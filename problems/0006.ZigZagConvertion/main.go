package leetcode

import (
	"fmt"
)

func convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || length == 0 {
		return s
	}
	groupCount := 2*numRows - 2
	group := length / groupCount
	mod := length % groupCount
	if mod > 0 {
		group = group + 1
	}

	var matrix [][]byte
	matrix = make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, group*2)
	}
	printMatrix(matrix)

	row, col := 0, 0
	for i, _ := range s {
		if i == 0 {
			row, col = 0, 0
		} else {
			mod := i % groupCount
			if mod == 0 {
				row = 0
				col++
			} else if mod < numRows {
				row++
			} else if mod == numRows {
				row--
				col++
			} else {
				row--
			}
		}
		matrix[row][col] = s[i]
	}
	printMatrix(matrix)
	return getMatrixString(matrix)
}

func getMatrixString(matrix [][]byte) string {
	var result []byte
	for i, rows := range matrix {
		for j, _ := range rows {
			if matrix[i][j] != byte(0) {
				result = append(result, matrix[i][j])
			}
		}
	}
	return string(result)
}

func printMatrix(matrix [][]byte) {
	for i, rows := range matrix {
		for j, _ := range rows {
			val := matrix[i][j]
			if val == byte(0) {
				val = byte(' ')
			}
			fmt.Print(string(val), ",")
		}
		fmt.Println("")
	}
}
