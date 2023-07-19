package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "prom_exmpl_http_duration_seconds",
		Help: "Duration of HTTP requests",
	}, []string{"path"})
)

func prometheusHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}


func main() {
	r := mux.NewRouter()
	r.Use(prometheusHandler)
	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/obj/{id}").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			time.Sleep(time.Second)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, World: %v\n", vars["id"])
		})


	http.ListenAndServe(":2112", r)
}
