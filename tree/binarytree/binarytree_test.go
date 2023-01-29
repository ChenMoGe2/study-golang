package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	n4 := NewNode(4, nil, nil)
	n5 := NewNode(5, nil, nil)
	n6 := NewNode(6, nil, nil)
	n2 := NewNode(2, n4, n5)
	n3 := NewNode(3, n6, nil)
	n1 := NewNode(1, n2, n3)

	fmt.Println("先序遍历")
	n1.Dlr(func(n *Node) {
		fmt.Println(n.data)
	})
	fmt.Println("先序遍历-非递归")
	n1.LDlr(func(n *Node) {
		fmt.Println(n.data)
	})
	fmt.Println("中序遍历")
	n1.Ldr(func(n *Node) {
		fmt.Println(n.data)
	})
	fmt.Println("中序遍历-非递归")
	n1.LLdr(func(n *Node) {
		fmt.Println(n.data)
	})
	fmt.Println("后序遍历")
	n1.Lrd(func(n *Node) {
		fmt.Println(n.data)
	})
	fmt.Println("后序遍历-非递归")
	n1.LLrd(func(n *Node) {
		fmt.Println(n.data)
	})
}
