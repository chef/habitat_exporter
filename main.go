package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listen_addr = flag.String("listen-address", ":9449",
		"The address/port to listen on for HTTP requests.")
)

func init() {
	flag.Parse()
	h := NewHabitatCollector()
	prometheus.MustRegister(h)
}

func main() {
	log.Println("Listening on", *listen_addr)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.RedirectHandler("/metrics", 302))
	log.Fatal(http.ListenAndServe(*listen_addr, nil))
}
