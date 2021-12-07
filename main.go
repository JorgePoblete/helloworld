package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	d, _ := time.ParseDuration("1s")
	time.Sleep(d)
	fmt.Fprintf(w, "Hello, world!!!\n")
	log.Printf("%s", r.URL)
}
