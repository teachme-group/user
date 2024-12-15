package client

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/teachme-group/user/pkg/errlist"
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

	encryptedPass, err := signer.EncryptPassword(request.Password)
	if err != nil {
		return response, err
	}

	if err = s.repos.SaveSignUpStep(ctx, request.SignUpToken, domain.SignUpStep{
		PrevStep: domain.StepStartSignUp,
		StepData: map[domain.Step]interface{}{
			domain.StepStartSignUp: domain.StartSignUpStepData{
				Email:              request.Email,
				Login:              request.Login,
				Password:           encryptedPass,
				SignUpSessionToken: request.SignUpToken,
				VerifyCode:         verifyCode,
			},
		},
	}, s.cfg.SignUpSessionTimeout); err != nil {
		return response, err
	}

	return &v1.SignUpInitResponse{
		SignUpToken:           request.GetSignUpToken(),
		SignUpSessionLifetime: s.cfg.SignUpSessionTimeout.Milliseconds(),
	}, nil
}

func (s *service) SignUpConfirm(
	ctx context.Context,
	request *v1.SignUpFinalizeRequest,
) (response *v1.SignUpFinalizeResponse, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	step, err := s.repos.GetSignUpStep(ctx, request.SignUpToken)
	if err != nil {
		return response, err
	}

	if step.PrevStep != domain.StepStartSignUp {
		return response, errlist.ErrInvalidSignUpStep
	}

	startData := domain.StartSignUpStepData{}

	err = mapstructure.Decode(step.StepData[domain.StepStartSignUp], &startData)
	if err != nil {
		return response, err
	}

	if startData.SignUpSessionToken != request.SignUpToken {
		return response, errlist.ErrInvalidSignUpToken
	}
	if startData.VerifyCode != request.VerificationCode {
		return response, errlist.ErrInvalidVerifyCode
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

	return &v1.SignUpFinalizeResponse{
		SessionId: resp.AccessToken,
	}, nil
}
