package util

import (
	"fmt"
	"math/rand"
	"time"
)

// simple log function for test only print value to terminal
func log(a ...interface{}) {
	fmt.Println(a...)
}

// swap is only visible within the package
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// random range in [min,max)
func random(min int, max int) int {
	// generate seed before everything happends
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
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

// testEqual return true if two slice is equal
func testEqual(a, b []interface{}) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
