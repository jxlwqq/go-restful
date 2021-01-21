package prometheus

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var duration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_duration_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, _ := mux.CurrentRoute(r).GetPathTemplate()
		timer := prometheus.NewTimer(duration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
