package errlist

import "github.com/Markuysa/pkg/errs"

var (
	ErrInvalidEmail              = errs.New("INVALID_EMAIL_FORMAT", errs.InvalidArgument, 10_1)
	ErrInvalidSignUpToken        = errs.New("INVALID_SIGN_UP_TOKEN", errs.InvalidArgument, 10_2)
	ErrInvalidSignUpStep         = errs.New("Invalid sign up step", errs.Internal, 10_3)
	ErrLoginOrEmailAlreadyExists = errs.New("INVALID_LOGIN_OR_EMAIL", errs.InvalidArgument, 10_4)
	ErrInvalidVerifyCode         = errs.New("INVALID_VERIFY_CODE", errs.InvalidArgument, 10_5)
	ErrProviderNotFound          = errs.New("INVALID_OAUTH_PROVIDER", errs.InvalidArgument, 10_6)
	ErrInvalidLoginCredentials   = errs.New("INVALID_EMAIL_OR_PASSWORD", errs.InvalidArgument, 10_7)
)
