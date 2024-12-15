package config

import (
	"github.com/Markuysa/pkg/consul"
	"github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/postgres"
	"github.com/Markuysa/pkg/prober"
	promLoager "github.com/Markuysa/pkg/prometheus"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
	"github.com/teachme-group/user/internal/misc/clients/session"
	"github.com/teachme-group/user/internal/service/client"
	"github.com/teachme-group/user/pkg/mail"
	"github.com/teachme-group/user/pkg/oauth"
)

type Config struct {
	GRPC         grpc.Config           `validate:"required" yaml:"grpc"`
	Postgres     postgres.PgxPoolCfg   `validate:"required" yaml:"postgres"`
	Redis        redis.Config          `validate:"required" yaml:"redis"`
	Logger       log.Config            `validate:"required" yaml:"logger"`
	SessionConn  session.Config        `validate:"required" yaml:"session_conn"`
	ClientConfig client.Config         `validate:"required" yaml:"client_config"`
	Mail         mail.Config           `validate:"required" yaml:"mail"`
	OauthConfig  oauth.ProvidersConfig `validate:"required" yaml:"oauth_config"`
	Consul       consul.Config         `validate:"required" yaml:"consul"`
	Probes       prober.Config         `validate:"required" yaml:"probes"`
	Prometheus   promLoager.Config     `validate:"required" yaml:"prometheus"`
}
