package matrix

import (
	"testing"

	"github.com/MarkWh1te/algo4go/util"
)

func createMatrix(t *testing.T) {
	util.RandomMatrix := util.RandomMatrix(3, 4)
	util.LogMatrix(util.RandomMatrix)
}

func testRotateMatrix(t *testing.T) {
	testMatrix := util.RandomMatrix(4, 4)
	util.Log("original matrix: ")
	util.LogMatrix(testMatrix)
	RotateMatrix(&testMatrix)
	util.Log("rotated matrix: ")
	util.LogMatrix(testMatrix)
}

// TestMatrix is main test for matrix
func TestMatrix(t *testing.T) {
	t.Run("test create matrix", createMatrix)
	t.Run("test rotate matrix", testRotateMatrix)
}
