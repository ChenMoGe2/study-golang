package sort

import (
	"fmt"
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{8, 5, 1, 10, 7, 3, 4}
	sort := &BubbleSort{}
	err := sort.Sort(array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(array)
}
