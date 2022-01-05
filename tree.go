package main

type Tree struct {
	head *Node
	len  int
}

type Node struct {
	key  string
	prev *Node
	next *Node
}

func (t *Tree) Insert(key string) {
	shifted := t.head
	t.head = &Node{key: key, prev: nil, next: shifted}
	t.len++
	if t.head.next != nil {
		t.head.next.prev = t.head
	}
}
