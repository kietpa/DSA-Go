package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func PrintList(root *Node) {
	// set first node
	nodeWalk := root

	// loop until node pointer is empty
	for nodeWalk.Next != nil {
		fmt.Println(nodeWalk.Value)
		nodeWalk = nodeWalk.Next // assign current node
	}
	fmt.Println(nodeWalk.Value)
}

func LenList(root *Node) int {
	nodeWalk := root
	count := 1
	for nodeWalk.Next != nil {
		count = count + 1
		nodeWalk = nodeWalk.Next
	}
	return count
}

func main() {
	// init linked list
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}
	aa.Next = &bb
	bb.Next = &cc

	PrintList(&aa)

	reversed := reverseList(&aa)

	PrintList(reversed)
}

func reverseList(head *Node) *Node {
	var prev, curr, front *Node
	curr = head

	for curr != nil {
		front = curr.Next // set front node to current node + 1
		curr.Next = prev  // set pointer at current node to point at previous node
		prev = curr       // move previous node to be at current node
		curr = front      // move current node to front node
	}
	return prev
	// iteration ends at
	// front node = nil
	// current node = nil
	// previous node = head
}

// article that helped me understand
// https://takeuforward.org/data-structure/reverse-a-linked-list/
