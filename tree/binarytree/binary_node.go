package binarytree

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  interface{}
}

func NewBinaryNode(data interface{}, left, right *BinaryNode) *BinaryNode {
	n := &BinaryNode{
		left:  left,
		right: right,
		data:  data,
	}
	return n
}

func (n *BinaryNode) SetLeft(node *BinaryNode) {
	if n == nil {
		return
	}
	n.left = node
}

func (n *BinaryNode) SetRight(node *BinaryNode) {
	if n == nil {
		return
	}
	n.right = node
}

//先序-递归
func (n *BinaryNode) Dlr(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	f(n)
	n.left.Dlr(f)
	n.right.Dlr(f)
}

//中序-递归
func (n *BinaryNode) Ldr(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	n.left.Ldr(f)
	f(n)
	n.right.Ldr(f)
}

//后序-递归
func (n *BinaryNode) Lrd(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	n.left.Lrd(f)
	n.right.Lrd(f)
	f(n)
}

//先序-非递归
func (n *BinaryNode) LDlr(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	cur := n
	rs := make([]*BinaryNode, 0)
	stack := make([]*BinaryNode, 0)
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
func (n *BinaryNode) LLdr(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	cur := n
	rs := make([]*BinaryNode, 0)
	stack := make([]*BinaryNode, 0)
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
func (n *BinaryNode) LLrd(f func(n *BinaryNode)) {
	if n == nil {
		return
	}
	cur := n
	var prev *BinaryNode
	rs := make([]*BinaryNode, 0)
	stack := make([]*BinaryNode, 0)
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
