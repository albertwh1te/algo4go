package algo4go

import (
	"github.com/MarkWh1te/algo4go/util"
)

// MaxHeapInsert keeps max heap and insert element into Maxheap
func MaxHeapInsert(arr *[]int, index int) {
	for (*arr)[index] > (*arr)[(index-1)/2] {
		util.Swap(&((*arr)[index]), &((*arr)[(index-1)/2]))
		index = (index - 1) / 2
	}
}

// MaxHeapify assumes that the binary trees rooted at LEFT and RIGHT are maxheaps
func MaxHeapify(arr *[]int, index int, size int) {
	var largest int
	left := index*2 + 1
	right := index*2 + 2

	for left < size {
		largest = index
		if (*arr)[left] > (*arr)[index] {
			largest = left
		}
		if right < size && (*arr)[right] > (*arr)[largest] {
			largest = right
		}
		// if largest is index means down up to righ place,termiate!
		if largest == index {
			return
		}
		util.Swap(&((*arr)[largest]), &((*arr)[index]))
		index = largest
		left = index*2 + 1
		right = index*2 + 2
	}
}

// HeapSort runs in O(nlog(n)) but cache poorly
func HeapSort(arr *[]int) {
	size := len(*arr)
	if len(*arr) < 2 {
		return
	}
	for i := 0; i < size; i++ {
		MaxHeapInsert(arr, i)
	}
	for n := size - 1; n >= 0; n-- {
		util.Swap(&((*arr)[0]), &((*arr)[n]))
		MaxHeapify(arr, 0, n)
	}
}
