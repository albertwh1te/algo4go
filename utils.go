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
