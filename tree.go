package main

import (
	"fmt"
	"strings"

	"github.com/influxdata/influxdb-client-go/api"
)

const (
	getTagKeys = "flux/tag_keys_by_measurement.flux"
	getValues  = "all_tag_values_by_tag.flux"
)

type Tree struct {
	Root *Node `json:"root"`
}

type Node struct {
	Key      string           `json:"key"`
	Parent   *Node            `json:"-"`
	Children map[string]*Node `json:"children"`
}

func (t *Tree) Insert(branch []string) {
	tokens := branch
	rootKey := tokens[0]
	tokens = tokens[1:]
	if t.Root == nil {
		t.Root = &Node{rootKey, nil, make(map[string]*Node)}
	} else if t.Root.Key != rootKey {
		panic(fmt.Sprintf("%v doesn't match %v", tokens, t.Root.Key))
	}

	m := t.Root.Children
	parent := t.Root
	for _, token := range tokens {
		if exists, ok := m[token]; ok {
			m = exists.Children // next iteration will look at the branch's Children
			continue
		}
		// t.tiers[i] = token
		m[token] = &Node{token, parent, map[string]*Node{}}
		m = m[token].Children
	}
}

func ruleToBranches(rule string, table *api.QueryTableResult) [][]string {
	tokens := strings.Split(rule, ",")
	var branches [][]string
	for table.Next() {
		record := table.Record()
		var branch []string
		for _, token := range tokens {
			if val, ok := record.Values()[token]; ok {
				branch = append(branch, val.(string))
				if token == "_field" {
					branch = append(branch, fmt.Sprintf("%.2f", record.Values()["_value"].(float64)))
				}
			}
		}
		branches = append(branches, branch)
	}
	return branches
}

func (t *Tree) Print() {
	if t.Root == nil {
		fmt.Println("empty-Tree")
		return
	}
	depth := 0
	fmt.Println(t.Root.String(depth))
}

func (n *Node) String(depth int) string {
	repr := ""
	repr = repr + fmt.Sprintf("depth=%d Key=%s\n", depth, n.Key)
	for _, child := range n.Children {
		repr = repr + strings.Repeat("\t", depth+1) + child.String(depth+1)
	}
	return repr
}
