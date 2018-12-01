package algo4go

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
			matrix[r][c] = random(1, 993)
		}
	}
	return matrix
}

//LogMatrix is a function print 2-D matrix line by line
func LogMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		// https://stackoverflow.com/a/31483763/5555392
		if i == 0 {
			log("[", m[i])
		} else if i == len(m)-1 {
			log(" ", m[i], "]")
		} else {
			log(" ", m[i])
		}
	}
}
