package Algorithm

import (
	"fmt"
	"math"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int, left *Node, right *Node) (*Node, error) {

	if data == 0 {
		return nil, fmt.Errorf("cannot pass zero values")
	}
	node := Node{
		Data:  data,
		Left:  left,
		Right: right,
	}
	return &node, nil
}

func (node *Node) Search(target int) *Node {
	if node.Data == target {
		fmt.Println("Found it !!")
		return node
	}
	if node.Left != nil && node.Data > target {
		return node.Left.Search(target)
	}
	if node.Right != nil && node.Data < target {
		return node.Right.Search(target)
	}
	return nil
}

func (node *Node) PreOrderTraversal() {
	fmt.Println(node.Data)
	if node.Left != nil {
		node.Left.PreOrderTraversal()
	}
	if node.Right != nil {
		node.Right.PreOrderTraversal()
	}
}

func (node *Node) PostOrderTraversal() {

	if node.Left != nil {
		node.Left.PostOrderTraversal()
	}
	if node.Right != nil {
		node.Right.PostOrderTraversal()
	}
	fmt.Println(node.Data)
}

func (node *Node) InOrderTraversal() {

	if node.Left != nil {
		node.Left.InOrderTraversal()
	}
	fmt.Println(node.Data)
	if node.Right != nil {
		node.Right.InOrderTraversal()
	}
}

func (node *Node) Height(h int) int {
	leftHeight := h
	if node.Left != nil {
		leftHeight = node.Left.Height(h + 1)
	}
	rightHeight := h
	if node.Right != nil {
		rightHeight = node.Right.Height(h + 1)
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight)))

}
