// Copyright 2018 Chef Software, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("listen-address", ":9449",
		"The address/port to listen on for HTTP requests.")
	habitatAddress = flag.String("habitat-address", "http://127.0.0.1:9631",
		"The address of the habitat supervisor API to query")
	versionFlag = flag.Bool("version", false,
		"Show the current version and exit")
)

func init() {
	flag.Parse()
	h := NewHabitatCollector(*habitatAddress)
	prometheus.MustRegister(h)
}

func main() {
	if *versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}
	log.Println("Listening on", *listenAddress)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.RedirectHandler("/metrics", 302))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
