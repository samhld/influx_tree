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
	root  *Node
	tree  *Tree
	tiers map[int]string
}

type Node struct {
	key      string
	parent   *Node
	children map[string]*Node
}

func (t *Tree) Insert(branch []string) {
	tokens := branch
	rootKey := tokens[0]
	tokens = tokens[1:]
	if t.root == nil {
		t.root = &Node{rootKey, nil, make(map[string]*Node)}
	} else if t.root.key != rootKey {
		panic(fmt.Sprintf("%v doesn't match %v", tokens, t.root.key))
	}

	m := t.root.children
	parent := t.root
	for _, token := range tokens {
		if exists, ok := m[token]; ok {
			m = exists.children // next iteration will look at the branch's children
			continue
		}
		// t.tiers[i] = token
		m[token] = &Node{token, parent, map[string]*Node{}}
		m = m[token].children
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
					branch = append(branch, fmt.Sprintf("%f", record.Values()["_value"].(float64)))
				}
			}
		}
		branches = append(branches, branch)
	}
	return branches
}

func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("empty-tree")
		return
	}
	depth := 0
	fmt.Println(t.root.String(depth))
}

func (n *Node) String(depth int) string {
	repr := ""
	repr = repr + fmt.Sprintf("depth=%d key=%s\n", depth, n.key)
	for _, child := range n.children {
		repr = repr + strings.Repeat("\t", depth+1) + child.String(depth+1)
	}
	return repr
}
