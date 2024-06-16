package application

import (
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var connectionCounter int64

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_duration_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var activeConnections = promauto.NewGaugeFunc(prometheus.GaugeOpts{
	Name: "http_active_connections",
	Help: "Current number of active connections.",
}, func() float64 {
	return float64(atomic.LoadInt64(&connectionCounter))
})

/*MetricsCheckHandler*/
func (goservice *GoService) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		atomic.AddInt64(&connectionCounter, 1)
		defer atomic.AddInt64(&connectionCounter, -1)
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		defer timer.ObserveDuration()

		next.ServeHTTP(w, r)
	})
}
