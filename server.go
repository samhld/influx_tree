package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/influxdata/influxdb-client-go/api"
)

type TreeServer struct {
	http.Handler
	tree  *Tree
	table *api.QueryTableResult
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error reading index.html: %v", err)
	}
	tmpl.Execute(w, nil)
}

func (ts *TreeServer) processRule(w http.ResponseWriter, r *http.Request) {
	rule := r.FormValue("rule")
	fmt.Printf("rule: %s\n", rule)
	tokens := strings.Split(rule, ",")
	ts.tree.setKeys(tokens)
	branches := ruleToBranches(tokens, ts.table)
	for _, branch := range branches {
		ts.tree.Insert(branch)
	}
	err := json.NewEncoder(w).Encode(ts.tree)
	if err != nil {
		fmt.Printf("error encoding JSON: %v\n", err)
	}
}
