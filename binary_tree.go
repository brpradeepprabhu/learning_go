package main

import (
	"fmt"
	"github.com/brpradeepprabhu/learning_go/Algorithm"
)

func main() {
	node, err := Algorithm.NewNode(10, nil, nil)
	if err != nil {
		fmt.Printf("Error in creating square %s", err)
	}
	left, _ := Algorithm.NewNode(5, nil, nil)
	if left != nil {
		node.Left = left
	}
	right, _ := Algorithm.NewNode(15, nil, nil)
	if left != nil {
		node.Right = right
	}

	left, _ = Algorithm.NewNode(2, nil, nil)
	if left != nil {
		node.Left.Left = left
	}
	right, _ = Algorithm.NewNode(6, nil, nil)
	if left != nil {
		node.Left.Right = right
	}

	left, _ = Algorithm.NewNode(13, nil, nil)
	if left != nil {
		node.Right.Left = left
	}
	right, _ = Algorithm.NewNode(1000, nil, nil)
	if left != nil {
		node.Right.Right = right
	}

	tree, err := Algorithm.NewTree(node, "Binary tree")
	if err != nil {
		fmt.Println("Error while creating tree", err)
	}

	searchTree := tree.Search(1000)
	fmt.Println(searchTree)
	tree.InOrderTraversal()

}
