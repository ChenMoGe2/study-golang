package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	n4 := NewBinaryNode(4, nil, nil)
	n5 := NewBinaryNode(5, nil, nil)
	n6 := NewBinaryNode(6, nil, nil)
	n2 := NewBinaryNode(2, n4, n5)
	n3 := NewBinaryNode(3, n6, nil)
	n1 := NewBinaryNode(1, n2, n3)

	fmt.Println("先序遍历")
	n1.Dlr(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
	fmt.Println("先序遍历-非递归")
	n1.LDlr(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
	fmt.Println("中序遍历")
	n1.Ldr(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
	fmt.Println("中序遍历-非递归")
	n1.LLdr(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
	fmt.Println("后序遍历")
	n1.Lrd(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
	fmt.Println("后序遍历-非递归")
	n1.LLrd(func(n *BinaryNode) {
		fmt.Println(n.data)
	})
}
