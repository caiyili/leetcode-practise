package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := x
	var reverse int
	for n > 0 {
		mod := n % 10
		reverse = reverse*10 + mod
		n = n / 10
	}
	return reverse == x
}
