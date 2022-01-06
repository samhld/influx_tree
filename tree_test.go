package main

import (
	"strings"
	"testing"
)

type Data map[string][]string

func TestTree(t *testing.T) {
	rule := "MEASUREMENT>region>host>FIELD"
	var tokens []string
	tokens = strings.Split(rule, ">")
}
