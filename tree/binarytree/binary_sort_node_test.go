package binarytree

import (
	"fmt"
	"testing"
)

func TestBinarySortTree(t *testing.T) {
	bst := NewBinarySortTree([]int{8, 3, 10, 1, 6, 7, 14, 4, 7, 13}, []interface{}{8, 3, 10, 1, 6, 7, 14, 4, 7, 13})
	scores, _ := bst.Sort()
	fmt.Println(scores)
}
