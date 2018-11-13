package search

import (
	"testing"
	"math/rand"
	"fmt"
)

func random(min int,max int) int{
	return rand.Intn(max-min) + min
}

// TestPartition generate test slice and check it's partition
func TestPartition(t *testing.T){
	const UPPER int = 100
	length := random(10, UPPER)
	// use slice beacse length is not const
	test := make([]int,length)
	for index := 0; index < length; index++ {
		test[index] = random(1,UPPER)
	}
	fmt.Println(test[length-1])
	partition(&test,0,length-1,length-1)
}
