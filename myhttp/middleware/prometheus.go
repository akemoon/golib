package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/akemoon/golib/myhttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func BaseMetrics() Midddleware {
	httpRequestsTotal := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"route", "status"},
	)

	httpRequestDuration := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 20, 30, 60},
		},
		[]string{"route"},
	)

	httpActiveRequests := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_active_requests",
			Help: "Number of active HTTP requests",
		},
		[]string{"route"},
	)

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()

				route := r.Pattern

				httpActiveRequests.WithLabelValues(route).Inc()

				wrap := myhttp.ResponseWriter{
					ResponseWriter: w,
					Status:         http.StatusOK,
				}

				defer func() {
					status := strconv.Itoa(wrap.Status)

					httpRequestsTotal.WithLabelValues(route, status).Inc()

					httpRequestDuration.WithLabelValues(route).Observe(time.Since(start).Seconds())
					httpActiveRequests.WithLabelValues(route).Dec()
				}()

				h.ServeHTTP(&wrap, r)
			},
		)
	}
}
