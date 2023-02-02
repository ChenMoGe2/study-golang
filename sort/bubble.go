package sort

type BubbleSort struct {
}

func (s *BubbleSort) Sort(items []int) error {
	if len(items) <= 1 {
		return nil
	}

	/*
		for i := 0; i < len(items) - 1; i++ {
			for j := i + 1; j < len(items); j++ {
				if items[j] < items[i] {
					temp := items[j]
					items[j] = items[i]
					items[i] = temp
				}
			}
		}
	*/

	maxIndex := 0
	minIndex := 0
	index := len(items) - 1

	for i := 0; i < len(items)-1; i++ {
		flag := true
		for j := 0; j < index; j++ {
			if items[j] > items[j+1] {
				temp := items[j]
				items[j] = items[j+1]
				items[j+1] = temp
				flag = false
				maxIndex = j
			}
		}
		if flag {
			break
		}

		index = maxIndex

		for j := index; j > minIndex; j-- {
			if items[j] < items[j-1] {
				temp := items[j]
				items[j] = items[j-1]
				items[j-1] = temp
				flag = false
			}
		}
		minIndex++
		if flag {
			break
		}
	}

	return nil
}
