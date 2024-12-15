package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/teachme-group/user/pkg/errlist"
	oauthCli "github.com/teachme-group/user/pkg/oauth"
	"github.com/teachme-group/user/pkg/random"
	"github.com/teachme-group/user/pkg/signer"
	"github.com/teachme-group/user/pkg/validate"

	sessionStorageV1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/teachme-group/user/internal/domain"
)

func (s *service) SignUpRequest(
	ctx context.Context,
	request *v1.SignUpInitRequest,
) (response *v1.SignUpInitResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	if !validate.Email(request.Email) {
		return response, errlist.ErrInvalidEmail
	}

	if err = s.repos.ValidateUserSignUp(ctx, request.Email); err != nil {
		return response, err
	}

	verifyCode := random.NewConfirmationCode()

	if err = s.mailer.Send(
		ctx,
		request.Email,
		"Sign Up",
		"Your code: "+verifyCode,
	); err != nil {
		return response, err
	}

	if err = s.repos.SaveSignUpStep(ctx, request.SignUpToken, domain.SignUpStep{
		PrevStep: domain.StepStartSignUp,
		StepData: domain.StepData{
			Email:              request.Email,
			Login:              request.Login,
			SignUpSessionToken: request.SignUpToken,
			VerifyCode:         verifyCode,
			SignUpSource:       domain.SignUpSourceNative,
		},
	}, s.cfg.SignUpSessionTimeout); err != nil {
		return response, err
	}

	return &v1.SignUpInitResponse{
		SignUpToken:           request.GetSignUpToken(),
		SignUpSessionLifetime: s.cfg.SignUpSessionTimeout.Milliseconds(),
	}, nil
}

func (s *service) SignUpConfirmEmail(
	ctx context.Context,
	request *v1.SignUpConfirmEmailRequest,
) (response *v1.SignUpConfirmEmailResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	step, err := s.repos.GetSignUpStep(ctx, request.SignUpToken)
	if err != nil {
		return response, err
	}

	if step.PrevStep != domain.StepStartSignUp {
		return response, errlist.ErrInvalidSignUpStep
	}

	startData := step.StepData

	if startData.SignUpSessionToken != request.SignUpToken {
		return response, errlist.ErrInvalidSignUpToken
	}
	if startData.VerifyCode != request.VerificationCode && startData.SignUpSource == domain.SignUpSourceNative {
		return response, errlist.ErrInvalidVerifyCode
	}

	step.StepData.EmailConfirmed = true

	if err = s.repos.SaveSignUpStep(ctx, request.SignUpToken, domain.SignUpStep{
		PrevStep: domain.ConfirmEmail,
		StepData: step.StepData,
	}, s.cfg.SignUpSessionTimeout); err != nil {
		return response, err
	}

	return &v1.SignUpConfirmEmailResponse{
		SignUpToken: request.SignUpToken,
	}, nil
}

func (s *service) SignUpEnterPassword(
	ctx context.Context,
	request *v1.SignUpEnterPasswordRequest,
) (response *v1.SignUpEnterPasswordResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	step, err := s.repos.GetSignUpStep(ctx, request.SignUpToken)
	if err != nil {
		return response, err
	}

	if step.PrevStep != domain.ConfirmEmail {
		return response, errlist.ErrInvalidSignUpStep
	}

	startData := step.StepData

	if startData.SignUpSessionToken != request.SignUpToken {
		return response, errlist.ErrInvalidSignUpToken
	}

	startData.Password, err = signer.EncryptPassword(request.Password)
	if err != nil {
		return response, err
	}

	user, err := s.repos.CreateUser(ctx, domain.User{
		Email:    startData.Email,
		Login:    startData.Login,
		Password: startData.Password,
	})
	if err != nil {
		return response, err
	}

	resp, err := s.sessionStorage.ClientSetSession(ctx, &sessionStorageV1.ClientSetSessionRequest{
		ClientId: user.ID.String(),
	})
	if err != nil {
		return response, err
	}

	return &v1.SignUpEnterPasswordResponse{
		SessionId: resp.AccessToken,
	}, nil
}

func (s *service) GetOauthSignUpUrls(
	ctx context.Context,
	in *v1.GetOauthSignUpUrlRequest,
) (*v1.GetOauthSignUpUrlResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	urls, err := s.oauth.AuthCodeURLs(uuid.NewString(), in.OauthProvider)
	if err != nil {
		return nil, err
	}

	return &v1.GetOauthSignUpUrlResponse{
		Urls: urls,
	}, nil
}

func (s *service) GoogleSignUpCallback(
	ctx context.Context,
	request *v1.HandleOauthCallbackRequest,
) (response *v1.HandleOauthCallbackResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	userInfo, err := s.oauth.ProcessCallback(
		ctx,
		oauthCli.Google,
		request.GetBody(),
		request.GetCallbackUrl(),
	)
	if err != nil {
		return response, err
	}

	sessionID := uuid.NewString()

	if err = s.repos.SaveSignUpStep(ctx, sessionID, domain.SignUpStep{
		PrevStep: domain.ConfirmEmail,
		StepData: domain.StepData{
			Email:              userInfo.Email,
			Login:              userInfo.Login,
			SignUpSource:       domain.SignUpSourceOauth,
			SignUpSessionToken: sessionID,
			EmailConfirmed:     true,
		},
	}, s.cfg.SignUpSessionTimeout); err != nil {
		return response, err
	}

	return &v1.HandleOauthCallbackResponse{
		SessionId: sessionID,
	}, nil
}
