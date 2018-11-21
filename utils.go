package algo4go

import (
	"fmt"
	"math/rand"
)

// swap is only visible within the package
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// random range open interval
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// get the max int in a slice
func maxIntSlice(v []int) int {
	if len(v) == 0 {
		return 0
	}
	max := v[0]
	for _, n := range v {
		if max < n {
			max = n
		}
	}
	return max
}

// generate random slice
func randomSlice(minlength int, maxlength int) []int {
	length := random(minlength, maxlength)
	// use slice because length is not a constant
	result := make([]int, length)
	for index := 0; index < length; index++ {
		// use 9973 because is the biggest number below 10000
		result[index] = random(1, 9973)
	}
	return result
}

// simple log function for test only print value to termial
func log(a ...interface{}) {
	fmt.Println(a...)
}
