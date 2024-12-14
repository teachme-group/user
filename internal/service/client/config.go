package client

import "time"

type Config struct {
	SIgnUpSessionTimeout time.Duration `envconfig:"SIGN_UP_SESSION_TIMEOUT" default:"10m"`
}
