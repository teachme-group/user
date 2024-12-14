package mail

type Config struct {
	SmtpHost string `validate:"required"`
	SmtpPort int    `validate:"required"`

	From     string `validate:"required"`
	Password string `validate:"required"`
}
