package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {

	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val: %d\n", n.val)
	}
}
