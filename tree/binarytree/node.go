package binarytree

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	tree := &BinaryTree{root: root}
	return tree
}

type Node struct {
	left  *Node
	right *Node
	data  interface{}
}

func NewNode(data interface{}, left, right *Node) *Node {
	n := &Node{
		left:  left,
		right: right,
		data:  data,
	}
	return n
}

//先序-递归
func (n *Node) Dlr(f func(n *Node)) {
	if n == nil {
		return
	}
	f(n)
	n.left.Dlr(f)
	n.right.Dlr(f)
}

//中序-递归
func (n *Node) Ldr(f func(n *Node)) {
	if n == nil {
		return
	}
	n.left.Ldr(f)
	f(n)
	n.right.Ldr(f)
}

//后序-递归
func (n *Node) Lrd(f func(n *Node)) {
	if n == nil {
		return
	}
	n.left.Lrd(f)
	n.right.Lrd(f)
	f(n)
}

//先序-非递归
func (n *Node) LDlr(f func(n *Node)) {
	if n == nil {
		return
	}
	cur := n
	rs := make([]*Node, 0)
	stack := make([]*Node, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			rs = append(rs, cur)
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
	}
	for _, r := range rs {
		f(r)
	}
}

//中序-非递归
func (n *Node) LLdr(f func(n *Node)) {
	if n == nil {
		return
	}
	cur := n
	rs := make([]*Node, 0)
	stack := make([]*Node, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rs = append(rs, cur)
		cur = cur.right
	}
	for _, r := range rs {
		f(r)
	}
}

//后序-非递归
func (n *Node) LLrd(f func(n *Node)) {
	if n == nil {
		return
	}
	cur := n
	var prev *Node
	rs := make([]*Node, 0)
	stack := make([]*Node, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.right == nil || cur.right == prev {
			rs = append(rs, cur)
			prev = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.right
		}
	}
	for _, r := range rs {
		f(r)
	}
}
