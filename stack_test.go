package algo4go

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	if s.Pop() != nil {
		t.Error("empty error")
	}
}
