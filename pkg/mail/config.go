package mail

type Config struct {
	SmtpHost string `validate:"required" yaml:"smtp_host"`
	SmtpPort int    `validate:"required" yaml:"smtp_port"`

	Username string `validate:"required"`
	Password string `validate:"required"`
}
