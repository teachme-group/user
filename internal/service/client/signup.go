package client

import (
	"context"
	"pkg/errlist"
	"pkg/signer"
	"pkg/validate"
	v1 "github.com/teachme-group/user/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/teachme-group/user/internal/domain"
)

func (s *service) SignUpRequest(
	ctx context.Context,
	request *v1.SignUpInitRequest,
) (signUpToken string, codeTTL int64, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	if !validate.Email(request.Email) {
		return signUpToken, codeTTL, errlist.ErrInvalidEmail
	}

	if err = s.repos.ValidateUserSignUp(ctx, request.Email); err != nil {
		return signUpToken, codeTTL, err
	}

	encryptedPass, err := signer.EncryptPassword(request.Password)
	if err != nil {
		return signUpToken, codeTTL, err
	}

	if err = s.repos.SaveSignUpStep(ctx, request.Email, domain.SignUpStep{
		PrevStep: domain.StepStartSignUp,
		StepData: map[domain.Step]interface{}{
			domain.StepStartSignUp: domain.StartSignUpStepData{
				Email:              request.Email,
				Login:              request.Login,
				Password:           encryptedPass,
				SignUpSessionToken: request.SignUpToken,
			},
		},
	}, s.cfg.SIgnUpSessionTimeout); err != nil {
		return signUpToken, codeTTL, err
	}

	return request.SignUpToken, 0, nil
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

	startData, ok := step.StepData[domain.StepStartSignUp].(domain.StartSignUpStepData)
	if !ok {
		return response, errlist.ErrInvalidSignUpStep
	}

	if startData.SignUpSessionToken != request.SignUpToken {
		return response, errlist.ErrInvalidSignUpToken
	}

	user, err := s.repos.CreateUser(ctx, domain.User{
		Email:    startData.Email,
		Login:    startData.Login,
		Password: startData.Password,
	})
	if err != nil {
		return response, err
	}

	if err

	return nil
}
