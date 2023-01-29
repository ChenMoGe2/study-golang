package binarytree

type BinarySortNode struct {
	left  *BinarySortNode
	right *BinarySortNode
	score int
	data  interface{}
}

func NewBinarySortTree(scores []int, datas []interface{}) *BinarySortNode {
	if len(scores) == 0 {
		return nil
	} else if len(scores) == 1 {
		n := &BinarySortNode{left: nil, right: nil, score: scores[0], data: datas[0]}
		return n
	} else {
		root := &BinarySortNode{left: nil, right: nil, score: scores[0], data: datas[0]}
		for i := 1; i < len(scores); i = i + 1 {
			score := scores[i]
			var data interface{}
			if len(datas) > i {
				data = datas[i]
			}
			curNode := root
			for {
				if score < curNode.score {
					if curNode.left == nil {
						node := &BinarySortNode{left: nil, right: nil, score: score, data: data}
						curNode.left = node
						break
					} else {
						curNode = curNode.left
					}
				} else {
					if curNode.right == nil {
						node := &BinarySortNode{left: nil, right: nil, score: score, data: data}
						curNode.right = node
						break
					} else {
						curNode = curNode.right
					}
				}
			}
		}
		return root
	}
}

func (b *BinarySortNode) Sort() ([]int, []interface{}) {
	if b == nil {
		return []int{}, []interface{}{}
	}
	cur := b
	scores := make([]int, 0)
	datas := make([]interface{}, 0)
	stack := make([]*BinarySortNode, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		scores = append(scores, cur.score)
		datas = append(datas, cur.data)
		cur = cur.right
	}
	return scores, datas
}
