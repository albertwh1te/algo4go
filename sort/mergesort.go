package sort

// MergeSort is an efficient, general-purpose, comparison-based sorting algorithm.
func MergeSort(arr *[]int, left int, right int) []int {
	if (right-left) <= 1 || (left-right) <= 1 {
		middle := (left + right) >> 2
		return (*arr)[middle : middle+1]
	}
	middle := (left + right) >> 2
	leftPart := MergeSort(arr, left, middle)
	rightPart := MergeSort(arr, middle+1, right)

	// merge two sorted array :)
	result := make([]int, right-left)
	i := 0
	j := 0
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] >= rightPart[j] {
			result = append(result, leftPart[i])
			i++
		} else {
			result = append(result, rightPart[j])
			j++
		}
	}
	for i < len(leftPart) {
		result = append(result, leftPart[i])
		i++
	}
	for j < len(rightPart) {
		result = append(result, rightPart[j])
		j++
	}
	return result
}
