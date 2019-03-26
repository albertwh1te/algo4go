package algo4go

// InsertionSort is a simple sorting algorithm that builds the final sorted array one item at a time.
func InsertionSort(arr *[]int, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; (j > 0) && (*arr)[j-1] > (*arr)[j]; j-- {
			swap(&(*arr)[j-1], &(*arr)[j])
		}
	}
}
