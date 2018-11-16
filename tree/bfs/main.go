package main

import (
	. "github.com/wwgberlin/algorithms/tree"
	"fmt"
)

var special = &Node{}

func Print(node *Node) {
	queue := []*Node{node}

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if node != nil {
			fmt.Print(node.Value, " ")
			queue = append(queue, node.Left, node.Right)
		}
	}
}

func main() {
	root := Node{Value: 8}
	n1 := Node{Value: 5}
	n2 := Node{Value: 10}
	n3 := Node{Value: 1}
	n4 := Node{Value: 6}
	n5 := Node{Value: 2}

	root.Left = &n1
	root.Right = &n3

	n1.Left = &n2
	n1.Right = &n4
	n3.Right = &n5

	fmt.Println(PrintTree(&root))
	Print(&root)
}
