package main

import "github.com/influxdata/influxdb-client-go/api"

type Tree struct {
	root   *Node
	leaves []*Node
}

type Node struct {
	key       string
	ancestors []*Node // parent and all levels above for filtering purposes
	children  map[string]*Node
}

func (t *Tree) Insert(key string, qAPI *api.QueryAPI) {
	if key == "MEASUREMENT" {
		qAPI.Query(getTagKeys)
		// TODO
	}
}

type Data map[string][]string

func (d Data) getKeyValues(key) {
	return d[key]
}

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
