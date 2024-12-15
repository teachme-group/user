package mail

import (
	"context"

	"github.com/Markuysa/pkg/tracer"
	"gopkg.in/gomail.v2"
)

type (
	mailer struct {
		cfg    Config
		dialer *gomail.Dialer
	}
)

func New(cfg Config) *mailer {
	return &mailer{
		cfg: cfg,
		dialer: gomail.NewDialer(
			cfg.SmtpHost,
			cfg.SmtpPort,
			cfg.Username,
			cfg.Password,
		),
	}
}

// TODO create preety pattern with html page
func (m *mailer) Send(ctx context.Context, to, subject, body string) error {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	msg := gomail.NewMessage()

	msg.SetBody("text/html", body)
	msg.SetHeader("Username", m.cfg.Username)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("From", m.cfg.Username)

	return m.dialer.DialAndSend(
		msg,
	)
}
