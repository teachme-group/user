package errlist

import "github.com/Markuysa/pkg/errs"

var (
	ErrInvalidEmail       = errs.New("Invalid email format", errs.InvalidArgument, 10_1)
	ErrInvalidSignUpToken = errs.New("Invalid sign up token", errs.InvalidArgument, 10_2)
	ErrInvalidSignUpStep  = errs.New("Invalid sign up step", errs.Internal, 10_3)
)
