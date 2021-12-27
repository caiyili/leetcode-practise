package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	result := make([]string, 0)
	for leftCount := 0; leftCount < n; leftCount++ {
		for _, left := range generateParenthesis(leftCount) {
			for _, right := range generateParenthesis(n - leftCount - 1) {
				result = append(result, "("+left+")"+right)
			}
		}
	}
	return result
}
