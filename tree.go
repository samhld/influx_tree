package main

import (
	"strings"

	"github.com/influxdata/influxdb-client-go/api"
)

const (
	getTagKeys = "flux/tag_keys_by_measurement.flux"
	getValues  = "all_tag_values_by_tag.flux"
)

type Node interface {
	Key() string
	Insert(key string, tierTokens []string, api *MeasurementAPI, ancestors []*Node)
}

type Tree struct {
	root  *Node
	tree  *Tree
	tiers map[int]string
}

type Key struct {
	key       string
	ancestors []*Node // parent and all levels above for filtering purposes
	children  map[string]*Node
}

type Value struct {
	key       string
	ancestors []*Node
	child     *Node
}

func (tt *Tree) Insert(tokens []string) {}

func ruleToBranches(rule string, table *api.QueryTableResult) [][]string {
	tokens := strings.Split(rule, ",")
	var branches [][]string
	for table.Next() {
		record := table.Record()
		var branch []string
		for _, token := range tokens {
			if val, ok := record.Values()[token]; ok {
				branch = append(branch, val.(string))
			}
		}
		branches = append(branches, branch)
	}
	return branches
}

// func createChildren(key string, vals []string) map[string]*Node {
// 	children := make(map[string]*Node)
// 	for _, val := range vals { // these values become keys to new nodes
// 		children[key] = NewNode(key, )
// 	}
// }
// func NewTree(rule string) Tree {
// 	for _, token := range tokens {
// 		switch token {
// 		case "MEASUREMENT":

// 		case "FIELD":

// 		default:
// 		}
// 	}
// }

// func (n *Node) String() string {
// 	key := fmt.Sprintf("key: %s\n", n.key)
// 	prev := ""
// 	next := ""
// 	if n.prev != nil {
// 		prev = fmt.Sprintf("prev: %q\n", n.prev)
// 	}
// 	if n.next != nil {
// 		next = fmt.Sprintf("next: %q\n", n.next)
// 	}
// 	return key + prev + next
// }
