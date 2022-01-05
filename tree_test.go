package main

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	measurement := "test"
	tree := Tree{}
	tree.Insert(measurement)
	tree.Insert("newNode")
	wantTree := Tree{
		head: &Node{
			key:  "newNode",
			prev: nil,
			next: &Node{
				key:  "test",
				prev: tree.head,
				next: nil,
			},
		},
		len: 2,
	}
	if !reflect.DeepEqual(tree, wantTree) {
		t.Errorf("\ngot %#v\nwant %#v\n", tree.head.key, wantTree.head.key)
	}
}
