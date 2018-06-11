package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("listen-address", ":9449",
		"The address/port to listen on for HTTP requests.")
	habitatAddress = flag.String("habitat-address", "http://127.0.0.1:9631",
		"The address of the habitat supervisor API to query")
)

func init() {
	flag.Parse()
	h := NewHabitatCollector(*habitatAddress)
	prometheus.MustRegister(h)
}

func main() {
	log.Println("Listening on", *listenAddress)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.RedirectHandler("/metrics", 302))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
