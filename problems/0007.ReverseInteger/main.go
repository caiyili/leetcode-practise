package leetcode

func reverse(x int) int {
	var result int
	flag := 1
	if x < 0 {
		flag = -1
		x = flag * x
	}
	for x > 0 {
		r := x % 10
		result = result*10 + r
		x = x / 10
	}
	result = result * flag
	if result < -(1<<31) || result > ((1<<31)-1) {
		return 0
	}
	return result
}
