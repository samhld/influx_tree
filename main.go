package main

import "fmt"

func main() {
	tree := Tree{}
	tree.Insert("test")
	tree.Insert("key1")
	tree.Insert("key2")
	printHeads(tree.head)
}

func printHeads(start *Node) {
	fmt.Printf("node: %s\n", start.key)
	fmt.Printf("node prev: %+v\n", start.prev)
	fmt.Printf("node next: %+v\n", start.next)
	if start.next != nil {
		printHeads(start.next)
	}
}
