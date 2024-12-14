package domain

type Step string

const (
	StepStartSignUp = Step("start_signup")
)

type (
	SignUpStep struct {
		PrevStep Step
		StepData map[Step]interface{}
	}

	StartSignUpStepData struct {
		Email              string
		Login              string
		Password           string
		SignUpSessionToken string
		VerifyCode         string
	}
)
