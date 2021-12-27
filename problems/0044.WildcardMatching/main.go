package leetcode

func isMatch(s string, p string) bool {
	p = compress(p)
	if s == "" {
		return p == "*" || p == ""
	} else if p == "" {
		return false
	}
	var lenStr, lenReg = len(s), len(p)
	var dp [][]bool = make([][]bool, lenStr)
	for i := 0; i < lenStr; i++ {
		dp[i] = make([]bool, lenReg)
	}

	var calcDpItem = func(i, j int) bool {
		if p[j] == '?' || p[j] == s[i] {
			if i > 0 && j > 0 {
				return dp[i-1][j-1]
			} else if j > 0 {
				return j == 1 && p[0] == '*'
			} else {
				return i == 0
			}
		} else if p[j] == '*' {
			if i > 0 && j > 0 {
				return dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]
			} else if j > 0 {
				return dp[i][j-1]
			} else {
				return true
			}
		} else {
			return false
		}
	}

	for i := 0; i < lenStr; i++ {
		for j := 0; j < lenReg; j++ {
			dp[i][j] = calcDpItem(i, j)
		}
	}

	return dp[lenStr-1][lenReg-1]
}

func compress(p string) string {
	buff := make([]byte, 0)
	for i := 0; i < len(p); i++ {
		if i > 0 && p[i] == '*' && p[i-1] == '*' {
			continue
		}
		buff = append(buff, p[i])
	}
	return string(buff)
}
