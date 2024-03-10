package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Register basic metrics
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"method", "path", "status"},
	)
	prometheus.MustRegister(httpRequestsTotal)

	// Success endpoint
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
		fmt.Println("Success")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	// Simulated failure endpoint
	http.HandleFunc("/maybe-fail", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(2) == 0 { // Randomly succeed or fail
			httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		} else {
			httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "500").Inc()
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
	})

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server.
	http.ListenAndServe(":8080", nil)
}
