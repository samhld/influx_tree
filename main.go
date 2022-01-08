package main

import (
	"context"
	"fmt"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	measAPI := setup()
	flux := fmt.Sprintf(readFlux("flux/single_query.flux"), measAPI.bucket, "test")
	result, err := measAPI.queryAPI.Query(context.Background(), flux)
	if err != nil {
		log.Fatalf("query err: %q", err)
	}

	branches := ruleToBranches("MEASUREMENT,region,host,_field", result)
	tree := Tree{}
	for _, branch := range branches {
		tree.Insert(branch)
	}

	for _, branch := range branches {
		fmt.Println(branch)
	}
	tree.Print()
}

func setup() *MeasurementAPI {
	measurement := "temp"
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	measAPI := &MeasurementAPI{queryAPI, bucket, measurement}

	return measAPI
}
