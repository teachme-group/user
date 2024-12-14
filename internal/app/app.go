package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Markuysa/pkg/closer"
	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/postgres"
	"github.com/Markuysa/pkg/srv/grpc"
	"github.com/teachme-group/user/internal/config"
)

func Run(_ context.Context, cfg *config.Config) error {
	cl := closer.New()

	err := logger.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		return err
	}
	cl.AddCloser(pg.Close)

	grpc, err := grpc.NewServer(
		grpc.WithConfig(&cfg.GRPC),
		grpc.WithRegistes(),
	)
	if err != nil {
		return err
	}
	cl.AddCloser(grpc.GracefulStop)

	logger.Logger.Info("started app")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	return cl.Close()
}
