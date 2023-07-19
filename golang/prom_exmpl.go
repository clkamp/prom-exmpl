package main

import (
	"fmt"
	"time"
	"net/http"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	r := mux.NewRouter()
	
	metrics := muxprom.NewDefaultInstrumentation()
	r.Use(metrics.Middleware)

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
