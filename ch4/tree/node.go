package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.right = &treeNode{}
	root.left = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root)

	nodes := []treeNode{{value: 3}, {5, nil, nil}, {}}
	fmt.Println(nodes)
}
