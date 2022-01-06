package main

import (
	"fmt"
	"os"
)

func readFlux(fileName string) string {
	byteFlux, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file: %s, err: %v", fileName, err)
	}
	return string(byteFlux)
}
