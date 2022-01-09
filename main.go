package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	measAPI := setup()
	flux := fmt.Sprintf(readFlux("flux/single_query.flux"), measAPI.bucket, measAPI.measurement)
	table, err := measAPI.queryAPI.Query(context.Background(), flux)
	if err != nil {
		log.Fatalf("query err: %q", err)
	}

	sMux := http.NewServeMux()
	tree := &Tree{}
	treeSvr := TreeServer{sMux, tree, table}
	sMux.HandleFunc("/process", treeSvr.processRule)
	sMux.HandleFunc("/", index)

	http.ListenAndServe(":5000", treeSvr)
}

func setup() *MeasurementAPI {
	measurement := "sensor"
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	measAPI := &MeasurementAPI{queryAPI, bucket, measurement}

	return measAPI
}
