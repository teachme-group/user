package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Markuysa/pkg/closer"
	"github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/postgres"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
	"github.com/teachme-group/user/internal/config"
	"github.com/teachme-group/user/internal/misc/clients/session"
	clientRepos "github.com/teachme-group/user/internal/repository/client"
	clientSrv "github.com/teachme-group/user/internal/service/client"
	clientTransp "github.com/teachme-group/user/internal/transport/client/grpc"
	"github.com/teachme-group/user/migration"
	"github.com/teachme-group/user/pkg/mail"
	"github.com/teachme-group/user/pkg/oauth"
)

func Run(_ context.Context, cfg *config.Config) error {
	cl := closer.New()

	pg, err := postgres.New(
		cfg.Postgres,
		postgres.WithMigrate(
			&postgres.MigrateCfg{
				MigratePath: ".",
				Fs:          migration.Migrations,
			},
		),
	)
	if err != nil {
		return err
	}
	cl.AddCloser(pg.Close)
	log.Info("postgres connected")

	rdConn, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}
	cl.AddErrCloser(rdConn.Close)
	log.Info("redis connected")

	sessionConn, err := session.NewClient(cfg.SessionConn)
	if err != nil {
		return err
	}
	cl.AddErrCloser(rdConn.Close)
	log.Info("connected to session service")

	mailer := mail.New(cfg.Mail)
	oauthClient := oauth.New(cfg.OauthConfig)

	clRepos := clientRepos.New(pg, rdConn)
	clService := clientSrv.New(
		cfg.ClientConfig,
		clRepos,
		sessionConn,
		mailer,
		oauthClient,
	)
	clTransport := clientTransp.New(clService)

	grpcSrv, err := grpc.NewServer(
		grpc.WithConfig(&cfg.GRPC),
		grpc.WithRegistes(
			clTransport,
		),
	)
	if err != nil {
		return err
	}
	log.Infof("grpc server created on %s", cfg.GRPC.Host)
	cl.AddCloser(grpcSrv.GracefulStop)

	log.Info("started app")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	return cl.Close()
}
