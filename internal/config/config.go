package config

import (
	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/postgres"
	"github.com/Markuysa/pkg/srv/grpc"
)

type Config struct {
	GRPC     grpc.Config
	Postgres postgres.PgxPoolCfg
	Logger   logger.Config
}
