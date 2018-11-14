package search

import (
	"fmt"
	"math/rand"
	"testing"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// TestPartition generate test slice and check it's partition
func TestPartition(t *testing.T) {
	const UPPER int = 1000
	length := random(10, UPPER)
	// use slice beacse length is not const
	test := make([]int, length)
	for index := 0; index < length; index++ {
		test[index] = random(1, UPPER)
	}
	pivot := test[length-1]
	fmt.Println("pivot: ", pivot)
	fmt.Println("origin arr: ", test)
	start, end := Partition(&test, 0, length-1, length-1)
	fmt.Println("start index: ", start, " end index: ", end)
	fmt.Println("new arr: ", test)
	for index := 0; index < length; index++ {
		if test[index] > pivot {
			if index < end {
				t.Errorf("big index check error")
			}
		} else if test[index] < pivot {
			if index > start {
				t.Errorf("small index check error")
			}
		}
	}
}
