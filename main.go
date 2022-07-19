package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/williamclot/emby_exporter/collector"
	"github.com/williamclot/emby_exporter/emby"
)

var (
	telemetryAddr = flag.String("telemetry.addr", ":9162", "address for the Emby exporter")
	metricsPath   = flag.String("telemetry.path", "/metrics", "URL path for surfacing collected metrics")

	embyAddr      = flag.String("emby.addr", ":8096", "address of Emby API")
	embyToken     = flag.String("emby.token", "", "Emby API key")
	embyVerifyTLS = flag.Bool("emby.verifyTLS", true, "verify TLS certificate of Emby Server")

	healthcheck = flag.Bool("health", false, "runs a small healthcheck against the exporter itself")
)

func main() {
	flag.Parse()

	if *healthcheck {
		c := http.Client{
			Timeout: time.Second,
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1%s/health", *telemetryAddr), nil)
		if err != nil {
			panic(err)
		}
		res, err := c.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			panic(fmt.Errorf("err %v, or unexpected status code %d", err, res.StatusCode))
		}
		os.Exit(0)
	}

	if *embyAddr == "" {
		log.Fatal("address of Emby server must be specified with '-emby.addr' flag")
	}
	if *embyToken == "" {
		log.Fatal("API key of Emby server must be specified with '-emby.key' flag")
	}

	ctx := context.Background()
	client := emby.New(ctx, *embyAddr, *embyToken, *embyVerifyTLS)

	// Creating prometheus registry and handler
	r := prometheus.NewRegistry()
	r.MustRegister(collector.New(client))
	mux := http.NewServeMux()
	mux.Handle(*metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Start listening for HTTP connections.
	log.Printf("starting Emby exporter on %q", *telemetryAddr)
	if err := http.ListenAndServe(*telemetryAddr, mux); err != nil {
		log.Fatalf("cannot start Emby exporter: %s", err)
	}
}
