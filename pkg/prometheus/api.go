package prometheus

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterHandlers(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler())
}
