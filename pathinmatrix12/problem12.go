package pathinmatrix12

func main() {}

func hasPath(matrix string, rows, cols int, str string) bool {
	strLen := len(str)
	if strLen == 0 {
		return false
	}
	matrixLen := len(matrix)
	if matrixLen == 0 {
		return false
	}
	if strLen > matrixLen {
		return false
	}
	if matrixLen != rows*cols {
		return false
	}

	runeMatrix := []rune(matrix)
	runeStr := []rune(str)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if p := check(i, j, rows, cols, runeMatrix, runeStr); p != nil {
				return true
			}
		}
	}
}

func check(i, j, rows, cols int, matrix []rune, str []rune) []int {
	if len(str) == 0 {
		return []int{}
	}
	if matrix[i*cols+j] != str[0] {
		return nil
	} else {
		if i-1 >= 0 {
			path := check(i-1, j, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if i+1 < rows {
			path := check(i+1, j, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if j-1 >= 0 {
			path := check(i, j-1, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if j+1 < cols {
			path := check(i, j+1, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		return nil
		}
	}
}

func existsAtPath(path []int, i, j int) bool {
	for p := 0; p < len(path); p += 2 {
		if path[p] == i && path[p+1] == j {
			return true
		}
	}
	return false
}
