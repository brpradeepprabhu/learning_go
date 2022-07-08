package main

import (
	"fmt"
	"github.com/brpradeepprabhu/learning_go/Algorithm"
)

func rotateRight(root *Algorithm.Node) *Algorithm.Node {
	pivot := root.Left
	reAttachNode := pivot.Right
	root.Left = reAttachNode
	pivot.Right = root
	return pivot
}

func rotateLeft(root *Algorithm.Node) *Algorithm.Node {
	pivot := root.Right
	reAttachNode := pivot.Left
	root.Right = reAttachNode
	pivot.Left = root
	return pivot
}
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
	if right != nil {
		node.Right = right
	}

	//second level children for 25
	left, _ = Algorithm.NewNode(10, nil, nil)
	if left != nil {
		node.Left.Left = left
	}
	right, _ = Algorithm.NewNode(35, nil, nil)
	if right != nil {
		node.Left.Right = right
	}
	//third level level children for 10
	left, _ = Algorithm.NewNode(5, nil, nil)
	if left != nil {
		node.Left.Left.Left = left
	}

	right, _ = Algorithm.NewNode(13, nil, nil)
	if right != nil {
		node.Left.Left.Right = right
	}

	//third level level children for 35

	left, _ = Algorithm.NewNode(30, nil, nil)
	if left != nil {
		node.Left.Right.Left = left
	}
	right, _ = Algorithm.NewNode(42, nil, nil)
	if right != nil {
		node.Left.Right.Right = right
	}

	//fourth level  children for 2
	left, _ = Algorithm.NewNode(2, nil, nil)
	if left != nil {
		node.Left.Left.Left.Left = left
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

	depthTree := tree.GetNodesAtDepth(3)

	fmt.Println("Depth at level2", depthTree)

	tree.AddNode(20)

	tree.DeleteNode(10)

	fmt.Println("Replaced Node", tree.Root.Left.Left)

	root, _ := Algorithm.NewNode(30, nil, nil)
	node20, _ := Algorithm.NewNode(20, nil, nil)
	node21, _ := Algorithm.NewNode(21, nil, nil)
	node10, _ := Algorithm.NewNode(10, nil, nil)
	node9, _ := Algorithm.NewNode(9, nil, nil)
	node11, _ := Algorithm.NewNode(11, nil, nil)

	rootRight, _ := Algorithm.NewNode(10, nil, nil)
	node20_1, _ := Algorithm.NewNode(20, nil, nil)
	node30, _ := Algorithm.NewNode(30, nil, nil)
	node19, _ := Algorithm.NewNode(19, nil, nil)
	node29, _ := Algorithm.NewNode(29, nil, nil)
	node31, _ := Algorithm.NewNode(31, nil, nil)

	if root != nil {
		unbalanceLeftLeft, _ := Algorithm.NewTree(root, "Unbalance Left Left")
		unbalanceLeftLeft.Root.Left = node20
		unbalanceLeftLeft.Root.Left.Right = node21
		unbalanceLeftLeft.Root.Left.Left = node10
		unbalanceLeftLeft.Root.Left.Left.Left = node9
		unbalanceLeftLeft.Root.Left.Left.Right = node11
		unbalanceLeftLeft.Rebalance()
		//unbalanceLeftLeft.Root = rotateRight(unbalanceLeftLeft.Root)
		fmt.Println("unbalanceRootLeft", unbalanceLeftLeft.Root.Data)
		fmt.Println("unbalanceRootLeft", unbalanceLeftLeft.Root.Left.Data)
		fmt.Println("unbalanceRootLeft", unbalanceLeftLeft.Root.Right.Data)

		unbalanceRightRight, _ := Algorithm.NewTree(rootRight, "Unbalance Right Right")
		unbalanceRightRight.Root.Right = node20_1
		unbalanceRightRight.Root.Right.Left = node19
		unbalanceRightRight.Root.Right.Right = node30
		unbalanceRightRight.Root.Right.Right.Left = node29
		unbalanceRightRight.Root.Right.Right.Right = node31
		unbalanceRightRight.Rebalance()
		//unbalanceRightRight.Root = rotateLeft(unbalanceRightRight.Root)
		fmt.Println("unbalanceRootRight", unbalanceRightRight.Root.Data)
		fmt.Println("unbalanceRootRight", unbalanceRightRight.Root.Left.Data)
		fmt.Println("unbalanceRootRight", unbalanceRightRight.Root.Left.Right.Data)
		fmt.Println("unbalanceRootRight", unbalanceRightRight.Root.Right.Data)
	}

}
