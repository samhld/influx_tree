package main

import (
	"fmt"
	"os"

	"github.com/influxdata/influxdb-client-go/api"
)

type MeasurementAPI struct {
	queryAPI    api.QueryAPI
	bucket      string
	measurement string
}

func readFlux(fileName string) string {
	byteFlux, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file: %s, err: %v", fileName, err)
	}
	return string(byteFlux)
}

// func genFlux(bucket, measurement string, filterCount int) string {
// 	s := `from(bucket: "%s")
// 	range(start: -6h)
// `
// 	for i < filterCount {
// 	}
// }

func genFilter(key, value string) string {
	return fmt.Sprintf(`  |> filter(fn: (r) => r.%s == "%s")`, key, value)
}
