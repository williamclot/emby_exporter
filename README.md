# Emby exporter


```
NAME:
  emby_exporter - A Prometheus exporter that exports metrics on Emby Media Server.

USAGE:
  emby_exporter [options]

OPTIONS:
  --telemetry.addr        Port for server (default: ":9162")
  --telemetry.path        URL path for metric collection (default: "/metrics")
  --emby.verifyTLS        Verify TLS certificate of Emby Server (default: true)
  --emby.addr             URL address of Emby API
  --emby.token            Emby API token
  --health                Run a healthcheck of the exporter
  --help, -h              Show help
```

