package matrix

import "github.com/MarkWh1te/algo4go/util"

// EmptyMatrix generate empty 2-D matrix
func EmptyMatrix(row int, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}
	return matrix
}

// RandomMatrix generate random numbers in matrix
func RandomMatrix(row int, col int) [][]int {
	matrix := EmptyMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			// use 10~99 for good print result
			matrix[r][c] = util.Random(10, 99)
		}
	}
	return matrix
}

// LogMatrix is a function print 2-D matrix line by line
func LogMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		if i == 0 {
			util.Log("[", m[i])
		} else if i == len(m)-1 {
			util.Log(" ", m[i], "]")
		} else {
			util.Log(" ", m[i])
		}
	}
}

// RotateMatrix clockwise rotation 90
func RotateMatrix(m *[][]int) {
	topRow := 0
	topCol := 0
	downRow := len(*m) - 1
	downCol := len((*m)[0]) - 1
	for topRow < downRow {
		rotateEdge(m, topRow, topCol, downRow, downCol)
		topRow++
		topCol++
		downRow--
		downCol--
	}
}

func rotateEdge(m *[][]int, topRow, topCol, downRow, downCol int) {
	times := downRow - topRow
	// save first value (topRow,topCOl)
	var tmp int
	for i := 0; i < times; i++ {
		tmp = (*m)[topRow][topCol+i]
		(*m)[topRow][topCol+i] = (*m)[downRow-i][topCol]
		(*m)[downRow-i][topCol] = (*m)[downRow][downCol-i]
		(*m)[downRow][downCol-i] = (*m)[topRow+i][downCol]
		(*m)[topRow+i][downCol] = tmp
	}
}
