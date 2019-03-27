package matrix

import (
	"testing"

	"github.com/MarkWh1te/algo4go/util"
)

func createMatrix(t *testing.T) {
	randomMatrix := RandomMatrix(3, 4)
	LogMatrix(randomMatrix)
}

func testRotateMatrix(t *testing.T) {
	testMatrix := RandomMatrix(4, 4)
	util.Log("original matrix: ")
	LogMatrix(testMatrix)
	RotateMatrix(&testMatrix)
	util.Log("rotated matrix: ")
	LogMatrix(testMatrix)
}

// TestMatrix is main test for matrix
func TestMatrix(t *testing.T) {
	t.Run("test create matrix", createMatrix)
	t.Run("test rotate matrix", testRotateMatrix)
}
