package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	tag    = "unknown"
	commit = "unknown"
)

var (
	runningVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "running_version",
		Help: "Running version of the application",
	},
		[]string{"version", "commit", "build_time"},
	)
)

func onBuild() {
	prometheus.MustRegister(runningVersion)

	builtAt := time.Now().String()

	runningVersion.WithLabelValues(tag, commit, builtAt).Set(1)
}
