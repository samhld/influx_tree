package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

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
	// if err != nil {
	// 	fmt.Printf("err reading form body: %v", err)
	// 	return
	// }
	branches := ruleToBranches(rule, ts.table)
	fmt.Printf("%q\n", branches)
	for _, branch := range branches {
		ts.tree.Insert(branch)
	}
	err := json.NewEncoder(w).Encode(ts.tree)
	if err != nil {
		fmt.Printf("error encoding JSON: %v\n", err)
	}
}
