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
	// flux := fmt.Sprintf(readFlux("flux/tag_keys_by_measurement.flux"), measAPI.bucket, measAPI.measurement)
	result, err := measAPI.queryAPI.Query(context.Background(), flux)
	if err != nil {
		log.Fatalf("query err: %q", err)
	}

	branches := ruleToBranches("MEASUREMENT,region,host,_field", result)
	fmt.Printf("branches: %q", branches)
}

func setup() *MeasurementAPI {
	measurement := "test"
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	measAPI := &MeasurementAPI{queryAPI, bucket, measurement}

	return measAPI
}

// func main() {
// 	tree := Tree{}
// 	tree.Insert("test")
// 	tree.Insert("key1")
// 	tree.Insert("key2")
// 	printHeads(tree.head)
// }

// func printHeads(start *Node) {
// 	fmt.Printf("node: %s\n", start.key)
// 	fmt.Printf("node prev: %+v\n", start.prev)
// 	fmt.Printf("node next: %+v\n", start.next)
// 	if start.next != nil {
// 		printHeads(start.next)
// 	}
// }
