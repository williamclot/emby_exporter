package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/williamclot/emby_exporter/collector"
	"github.com/williamclot/emby_exporter/emby"
)

var (
	telemetryAddr = flag.String("telemetry.addr", ":9162", "address for the emby exporter")
	metricsPath   = flag.String("telemetry.path", "/metrics", "URL path for surfacing collected metrics")

	embyAddr = flag.String("emby.addr", ":8096", "address of emby API")
	embyKey  = flag.String("emby.key", "", "emby API key")
)

func main() {
	flag.Parse()
	if *embyAddr == "" {
		log.Fatal("address of emby server must be specified with '-emby.addr' flag")
	}
	if *embyKey == "" {
		log.Fatal("API key of emby server must be specified with '-emby.key' flag")
	}

	ctx := context.Background()
	client := emby.New(ctx, *embyAddr, *embyKey)

	// Creating prometheus registry and handler
	r := prometheus.NewRegistry()
	r.MustRegister(collector.New(client))
	mux := http.NewServeMux()
	mux.Handle(*metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	// Start listening for HTTP connections.
	log.Printf("starting emby exporter on %q", *telemetryAddr)
	if err := http.ListenAndServe(*telemetryAddr, mux); err != nil {
		log.Fatalf("cannot start emby exporter: %s", err)
	}
}
