package main

import (
	"ch4/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Right = &tree.Node{}
	root.Left = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	fmt.Println(root)

	nodes := []tree.Node{{Value: 3}, {5, nil, nil}, {}}
	fmt.Println(nodes)

	root.Print()

	var newNode *tree.Node
	newNode.SetValue(200)
	newNode = &root
	newNode.SetValue(300)
	newNode.Print()
	root.Traverse()
}
