package main

// slo-workbench
// Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.
// Domain: monitoring

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqs = prometheus.NewCounterVec(prometheus.CounterOpts{
    Name: "requests_total", Help: "Total requests"}, []string{"endpoint"})

func main() {
    prometheus.MustRegister(reqs)
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"status":"ok","project":"slo-workbench","domain":"monitoring"})
    })
    http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
        // TODO: add real readiness check
        w.Write([]byte(`{"ready":true}`))
    })
    http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]interface{}{"firing":0,"silenced":0})
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        reqs.WithLabelValues("/").Inc()
        json.NewEncoder(w).Encode(map[string]string{"message":"Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.","service":"slo-workbench"})
    })
    log.Println("slo-workbench listening :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
