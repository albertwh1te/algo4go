package sort

import (
	"github.com/MarkWh1te/algo4go/util"
)

// InsertionSort is a simple sorting algorithm that builds the final sorted array one item at a time.
func InsertionSort(arr *[]int, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; (j > 0) && (*arr)[j-1] > (*arr)[j]; j-- {
			util.Swap(&(*arr)[j-1], &(*arr)[j])
		}
	}
}
