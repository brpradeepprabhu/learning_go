package main

import (
	"fmt"
	"github.com/brpradeepprabhu/learning_go/Algorithm"
)

func main() {
	//root node
	node, err := Algorithm.NewNode(50, nil, nil)

	if err != nil {
		fmt.Printf("Error in creating square %s", err)
	}

	//first children
	left, _ := Algorithm.NewNode(25, nil, nil)
	if left != nil {
		node.Left = left
	}
	right, _ := Algorithm.NewNode(75, nil, nil)
	if left != nil {
		node.Right = right
	}

	//second level children for 25
	left, _ = Algorithm.NewNode(10, nil, nil)
	if left != nil {
		node.Left.Left = left
	}
	right, _ = Algorithm.NewNode(35, nil, nil)
	if left != nil {
		node.Left.Right = right
	}
	//third level level children for 10
	left, _ = Algorithm.NewNode(5, nil, nil)
	if left != nil {
		node.Left.Left.Left = left
	}
	//fourth level  children for 2
	left, _ = Algorithm.NewNode(2, nil, nil)
	if left != nil {
		node.Left.Left.Left.Left = left
	}
	right, _ = Algorithm.NewNode(13, nil, nil)
	if left != nil {
		node.Left.Left.Right = right
	}
	//third level level children for 35

	left, _ = Algorithm.NewNode(30, nil, nil)
	if left != nil {
		node.Left.Right.Left = left
	}
	right, _ = Algorithm.NewNode(42, nil, nil)
	if left != nil {
		node.Left.Right.Right = right
	}
	tree, err := Algorithm.NewTree(node, "Binary tree")
	if err != nil {
		fmt.Println("Error while creating tree", err)
	}

	searchTree := tree.Search(1000)
	fmt.Println(searchTree)
	tree.InOrderTraversal()

	maxHeight := tree.FindMaxHeight()

	fmt.Println("Max Height", maxHeight)

}
