package testing

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestEncode(t *testing.T) {
	s := "hello，万志军"
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))

	fmt.Println("-------")
	for i, v := range s {
		t.Log(i, v)
		t.Log(string(v))
	}

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		ch := fmt.Sprintf("%c", r)
		t.Log(ch)
	}
}

func TestArr(t *testing.T) {
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := new([3]int)
	t.Logf("%T", arr1)
	t.Logf("%T", arr2)
}

type treeNode struct {
	value       string
	left, right *treeNode
}

func (node *treeNode) traverse() {
	fmt.Println(node.value)
	if node.left != nil {
		node.left.traverse()
	}
	if node.right != nil {
		node.right.traverse()
	}
}

func createTreeNode() treeNode {
	root := treeNode{"A", nil, nil}
	root.left = &treeNode{value: "B"}
	root.right = &treeNode{value: "C"}
	root.left.left = &treeNode{value: "D"}
	root.left.right = &treeNode{value: "E"}
	root.right.left = &treeNode{value: "F"}
	root.right.right = &treeNode{value: "G"}
	return root
}

func TestPrintNode(t *testing.T) {
	node := createTreeNode()
	node.traverse()
}
