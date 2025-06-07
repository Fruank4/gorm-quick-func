package Leetecode

func setZeroes(matrix [][]int) {

	row, col := false, false

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			row = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	high, low, left, right := -1, len(matrix), -1, len(matrix[0])

	ans := make([]int, 0)
	i, j := 0, 0
	for len(ans) < m*n {
		for j < right {
			ans = append(ans, matrix[i][j])
			j++
		}
		high++
		j--
		i++
		if len(ans) >= m*n {
			break
		}

		for i < low {
			ans = append(ans, matrix[i][j])
			i++
		}
		right--
		i--
		j--
		if len(ans) >= m*n {
			break
		}

		for j > left {
			ans = append(ans, matrix[i][j])
			j--
		}
		low--
		j++
		i--
		if len(ans) >= m*n {
			break
		}

		for i > high {
			ans = append(ans, matrix[i][j])
			i--
		}
		left++
		i++
		j++
		if len(ans) >= m*n {
			break
		}
	}

	return ans
}

func rotateMatrix(matrix [][]int) {
	n := len(matrix)
	// 1 1 1 1 1
	// 1 1 1 1 1
	// 1 1 1 1 1
	// 1 1 1 1 1
	// 1 1 1 1 1
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
