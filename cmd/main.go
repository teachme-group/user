package main

import (
	"context"
	"log"
	"os"
	"time"

	cfgLoader "github.com/Markuysa/pkg/config"
	"github.com/Markuysa/pkg/consul"
	logger "github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/prober"
	promLoager "github.com/Markuysa/pkg/prometheus"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/teachme-group/user/internal/app"
	"github.com/teachme-group/user/internal/config"
)

var (
	tag            = "unknown"
	commit         = "unknown"
	runningVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "running_version",
		Help: "Running version of the application",
	},
		[]string{"version", "commit", "build_time"},
	)
)

const (
	cfgPathKey = "CONFIG_PATH"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ctx := context.Background()

	cfgPath := os.Getenv(cfgPathKey)
	cfg := &config.Config{}

	err = cfgLoader.LoadFromYAML(cfg, cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := registerAnalytics(cfg); err != nil {
		log.Fatal(err)
	}

	if err = app.Run(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func registerAnalytics(cfg *config.Config) error {
	err := logger.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	prometheus.MustRegister(runningVersion)
	builtAt := time.Now().String()
	runningVersion.WithLabelValues(tag, commit, builtAt).Set(1)

	err = consul.RegisterService(cfg.Consul)
	if err != nil {
		return err
	}
	logger.Info("registered in consul")

	err = promLoager.LaunchPrometheusListener(cfg.Prometheus)
	if err != nil {
		return err
	}
	logger.Info("launched prometheus listener")

	err = prober.LaunchProbes(cfg.Probes)
	if err != nil {
		return err
	}
	logger.Info("launched probes")

	return nil
}
