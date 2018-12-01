package algo4go

import "testing"

func createMatrix(t *testing.T) {
	randomMatrix := RandomMatrix(3, 4)
	LogMatrix(randomMatrix)
}

// TestMatrix is main test for matrix
func TestMatrix(t *testing.T) {
	t.Run("test create matrix", createMatrix)
}
