package mail

import "context"

type (
	mailer struct {
		cfg Config
	}
)

func New(cfg Config) *mailer {
	return &mailer{cfg: cfg}
}

func (m *mailer) Send(ctx context.Context, to, subject, body string) error {
	return nil
}
