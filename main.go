package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthcheck", healthcheck)
	http.ListenAndServe(":8080", nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.URL)
	count := 1
	for i := 0; i <= 1000000; i++ {
		count = i
	}
	fmt.Fprintf(w, "Hello, world %d times\n", count)
}
