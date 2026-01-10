package middleware

import (
	"net/http"
	"strconv"

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

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				wrap := myhttp.ResponseWriter{
					ResponseWriter: w,
					Status:         http.StatusOK,
				}

				h.ServeHTTP(&wrap, r)

				route := r.Pattern
				status := strconv.Itoa(wrap.Status)

				httpRequestsTotal.WithLabelValues(route, status).Inc()
			},
		)
	}
}
