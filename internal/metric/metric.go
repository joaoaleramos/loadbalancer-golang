package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsHandled = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "loadbalancer_requests_total",
			Help: "Number of requests handled by each backend",
		},
		[]string{"backend"},
	)
)

func init() {
	prometheus.MustRegister(RequestsHandled)
}

// MetricsHandler creates an HTTP handler that exposes Prometheus metrics.
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
