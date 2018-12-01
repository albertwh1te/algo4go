package algo4go

import "testing"

func createMatrix(t *testing.T) {
	randomMatrix := RandomMatrix(3, 4)
	LogMatrix(randomMatrix)
}

func testRotateMatrix(t *testing.T) {
	testMatrix := RandomMatrix(4, 4)
	log("original matrix: ")
	LogMatrix(testMatrix)
	RotateMatrix(&testMatrix)
	log("rotated matrix: ")
	LogMatrix(testMatrix)
}

// TestMatrix is main test for matrix
func TestMatrix(t *testing.T) {
	t.Run("test create matrix", createMatrix)
	t.Run("test rotate matrix", testRotateMatrix)
}
