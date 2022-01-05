package main

import (
	"github.com/influxdata/influxdb-client-go/api"
)

type MeasurementAPI struct {
	api api.QueryAPI
}
