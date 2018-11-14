/*
Package search contains everything about search
*/
package search

// swap is only visible within the package
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// Partition moves the small number to head ,big number to end
func Partition(arr *[]int, l int, r int, n int) (int, int) {
	start := l
	current := l
	end := r
	pivot := (*arr)[n]
	for current <= end {
		if (*arr)[current] < pivot {
			swap(&(*arr)[current], &(*arr)[start])
			current++
			start++
		} else if (*arr)[current] > pivot {
			swap(&(*arr)[current], &(*arr)[end])
			end--
		} else if (*arr)[current] == pivot {
			current++
		}
	}
	return start, end
}

// QuickSort is a randomize sorting algorithms
func QuickSort(arr *[]int, l int, r int) {
}
