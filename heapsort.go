package algo4go

// MaxHeapInsert keeps max heap and insert element into Maxheap
func MaxHeapInsert(arr *[]int, index int) {
	for (*arr)[index] > (*arr)[(index-1)/2] {
		swap(&((*arr)[index]), &((*arr)[(index-1)/2]))
		index = (index - 1) / 2
	}
}

// MaxHeapify assumes that the binary trees rooted at LEFT and RIGHT are maxheaps
func MaxHeapify(arr *[]int, index int, size int) {
	left := index*2 + 1
	if left < size {
	}
}
