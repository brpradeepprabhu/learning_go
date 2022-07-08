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

func (node *Node) GetNodeAtDepth(depth int, nodes *[]int) []int {
	if depth == 0 {

		*nodes = append(*nodes, node.Data)
		fmt.Println("nodes", nodes)
	}
	if node.Left != nil {
		node.Left.GetNodeAtDepth(depth-1, nodes)
	} else {
		number := math.Pow(2.0, float64(depth-1))
		for i := 0; i < int(number); i++ {
			*nodes = append(*nodes, 0)
		}
	}
	if node.Right != nil {
		node.Right.GetNodeAtDepth(depth-1, nodes)
	} else {
		number := math.Pow(2.0, float64(depth-1))
		for i := 0; i < int(number); i++ {
			*nodes = append(*nodes, 0)
		}
	}
	return *nodes
}

func (node *Node) AddNode(data int) {
	if node.Data == data {
		panic("Tree cannot have the same data to be inserted")
	}
	if node.Data > data {
		if node.Left == nil {
			fmt.Println("inserting node", node.Data)
			left, _ := NewNode(data, nil, nil)
			if left != nil {
				node.Left = left
			}
		} else {
			node.Left.AddNode(data)
		}
	}
	if node.Data < data {

		if node.Right == nil {
			fmt.Println("inserting node", node.Data)
			right, _ := NewNode(data, nil, nil)
			if right != nil {
				node.Right = right
			}
		} else {
			node.Right.AddNode(data)
		}
	}
}
func (node *Node) findMin() int {
	if node.Left != nil {
		return node.Left.findMin()
	}
	return node.Data
}

func (node *Node) DeleteNode(target int) *Node {
	if node.Data == target {
		if node.Left != nil && node.Right != nil {
			// RTFM
			minValue := node.Right.findMin()
			//Replacing the value with minValue
			node.Data = minValue
			//Deleting the minvalue to avoid duplicate
			node.Right = node.Right.DeleteNode(minValue)

		} else {
			if node.Left != nil {
				return node.Left
			} else if node.Right != nil {
				return node.Right
			} else {
				return nil
			}
		}

	}
	if node.Right != nil && target > node.Data {
		node.Right = node.Right.DeleteNode(target)
	}
	if node.Left != nil && target < node.Data {
		node.Left = node.Left.DeleteNode(target)
	}
	return node
}

func (node *Node) IsBalanced() bool {
	leftHeight := node.Left.Height(0)
	rightHeight := node.Right.Height(0)
	return math.Abs(float64(leftHeight-rightHeight)) < 2
}
