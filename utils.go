package algo4go

import (
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

// get heap parent element
func parent(i int) int {
	return i - 1/2
}

// get heap left element
func left(i int) int {
	return 2*i + 1
}

// get heap right element
func right(i int) int {
	return 2*i + 2
}
