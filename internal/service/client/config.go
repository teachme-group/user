package client

import "time"

type (
	Config struct {
		SignUpSessionTimeout time.Duration `envconfig:"SIGN_UP_SESSION_TIMEOUT" default:"10m"`
	}
)
