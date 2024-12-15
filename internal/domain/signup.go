package domain

type Step string

const (
	StepStartSignUp = Step("start_signup")
	ConfirmEmail    = Step("confirm_email")
	EnterPassword   = Step("enter_password")
)

type (
	SignUpSource string

	SignUpStep struct {
		PrevStep Step
		StepData StepData
	}

	StepData struct {
		Email              string
		Login              string
		Password           string
		SignUpSessionToken string
		VerifyCode         string
		SignUpSource       SignUpSource
		EmailConfirmed     bool
	}
)

const (
	SignUpSourceNative = SignUpSource("native")
	SignUpSourceOauth  = SignUpSource("oauth")
)
