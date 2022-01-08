package main

// import (
// 	"reflect"
// 	"testing"
// )

// func TestTree(t *testing.T) {
// 	rule := "MEASUREMENT>region>host>FIELD"
// 	var tokens []string
// 	tokens = strings.Split(rule, ">")
// 	data := map[string][]string{
// 		"_measurement": {"test"},
// 		"region": {"us-west", "us-east"},
// 		"host": {"001", "002", "003"},
// 		"_field": {"usage_user"},
// 		"_value": {"45.0"},
// 	}
// 	head := &Node{"test", nil, nil}
// 	tree := Tree{head}

// 	t.Run("test create single node", func(t *testing.T) {
// 		vals := {"us-west", "us-east"}
// 		tree.Insert("region")
// 		wantTree := Tree{
// 			root: &Node{
// 				key: "test",
// 				ancestors: nil,
// 				children: []&Node{
// 					{key: "region", &Node{"test", nil}, children: nil},
// 				},
// 			},
// 			len: 2,
// 		}
// 	}
// }
