package main

import (
	"context"
	"log"
	"os"
	"time"

	cfgLoader "github.com/Markuysa/pkg/config"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/teachme-group/user/config"
	"github.com/teachme-group/user/internal/app"
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
	ctx := context.Background()
	onBuild()

	cfgPath := os.Getenv(cfgPathKey)
	cfg := &config.Config{}

	err := cfgLoader.LoadFromYAML(cfg, cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func onBuild() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	prometheus.MustRegister(runningVersion)

	builtAt := time.Now().String()

	runningVersion.WithLabelValues(tag, commit, builtAt).Set(1)
}
