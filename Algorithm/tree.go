package Algorithm

import (
	"fmt"
)

type Tree struct {
	Root *Node
	Name string
}

func NewTree(root *Node, name string) (*Tree, error) {

	if root == nil {
		return nil, fmt.Errorf("Not able to create the tree")
	}
	tree := Tree{
		Root: root,
		Name: name,
	}
	return &tree, nil
}

func (tree *Tree) Search(target int) *Node {
	return tree.Root.Search(target)
}

func (tree *Tree) PreOrderTraversal() {
	tree.Root.PreOrderTraversal()
}

func (tree *Tree) PostOrderTraversal() {
	tree.Root.PostOrderTraversal()
}

func (tree *Tree) InOrderTraversal() {
	tree.Root.InOrderTraversal()
}

func (tree *Tree) FindMaxHeight() int {
	return tree.Root.Height(0)
}

func (tree *Tree) GetNodesAtDepth(depth int) []int {
	nodes := []int{}
	return tree.Root.GetNodeAtDepth(depth, &nodes)
}

func (tree *Tree) AddNode(data int) {
	tree.Root.AddNode(data)
}

func (tree *Tree) DeleteNode(data int) {
	tree.Root = tree.Root.DeleteNode(data)
}

func (tree *Tree) Rebalance() {
	tree.Root.Rebalance()
	tree.Root = tree.Root.CheckImbalanceIfExists()
}
