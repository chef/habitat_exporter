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
	"fmt"
	"log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type habitatCollector struct {
	healthDesc     *prometheus.Desc
	habitatAddress string
}

type habitatService struct {
	ServiceGroup string `json:"service_group"`
}

const (
	habStatusOK = iota
	habStatusWarning
	habStatusCritical
	habStatusUnknown
)

// Constructor
func NewHabitatCollector(habitatAddress string) prometheus.Collector {
	return &habitatCollector{
		healthDesc: prometheus.NewDesc(
			"habitat_service_health",
			"Habitat Service Health",
			[]string{"service_group"},
			nil,
		),
		habitatAddress: habitatAddress,
	}
}

func (c *habitatCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.healthDesc
}

func (c *habitatCollector) Collect(ch chan<- prometheus.Metric) {
	services := []habitatService{}
	err := JsonHttpGet(fmt.Sprintf("%s/services", c.habitatAddress), &services)
	if err != nil {
		log.Println("Error getting list of services:", err)
		return
	}

	for _, service := range services {
		url := fmt.Sprintf("%s/services/%s/health", c.habitatAddress,
			strings.Replace(service.ServiceGroup, ".", "/", 1))
		httpStatus, err := HttpGetStatus(url)
		if err != nil {
			// Don't skip this - httpStatus will be 0 if there is an error, so
			// it will show up as "Unknown" in the prometheus output, which is
			// what we want
			log.Println("Error getting habitat status for",
				service.ServiceGroup, "-", err)
		}

		// Warning isn't implemented yet because it returns 200, and the
		// health_check value in services always returns "Unknown".
		// See https://github.com/habitat-sh/habitat/issues/4988
		value := habStatusUnknown
		if httpStatus == 200 {
			value = habStatusOK
		} else if httpStatus == 503 {
			value = habStatusCritical
		}

		ch <- prometheus.MustNewConstMetric(
			c.healthDesc,
			prometheus.GaugeValue,
			float64(value),
			service.ServiceGroup,
		)
	}
}
