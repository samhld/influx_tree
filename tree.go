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
	Root *Node    `json:"root"`
	Keys []string `json:"-"`
}

type Node struct {
	Key      string           `json:"key"`
	Parent   *Node            `json:"-"`
	Children map[string]*Node `json:"children"`
}

func (t *Tree) Insert(branch []string) {
	fmt.Printf("branch: %q\n", branch)
	rootKey := branch[0]
	branch = branch[1:]
	if t.Root == nil {
		t.Root = &Node{rootKey, nil, make(map[string]*Node)}
	} else if t.Root.Key != rootKey {
		panic(fmt.Sprintf("%v doesn't match %v", branch, t.Root.Key))
	}

	m := t.Root.Children
	parent := t.Root
	for _, bElem := range branch {
		if exists, ok := m[bElem]; ok {
			fmt.Printf("child %#v exists", m[bElem])
			m = exists.Children
			fmt.Printf("exists children: %#v\n", m)
			continue
		}
		m[bElem] = &Node{bElem, parent, map[string]*Node{}}
		m = m[bElem].Children
	}
}

func ruleToBranches(tokens []string, table *api.QueryTableResult) [][]string {
	var branches [][]string
	for table.Next() {
		record := table.Record()
		var branch []string
		for i, token := range tokens {
			if val, ok := record.Values()[token]; ok {
				if (i > 0) && (i < len(tokens)-1) {
					branch = append(branch, token)
				}
				branch = append(branch, val.(string))
				if token == "_field" {
					branch = append(branch, fmt.Sprintf("%.2f", record.Values()["_value"].(float64)))
				}
			}
		}
		fmt.Printf("branch: %q\n", branch)
		branches = append(branches, branch)
	}
	return branches
}

func (t *Tree) setKeys(tokens []string) {
	t.Keys = tokens
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
