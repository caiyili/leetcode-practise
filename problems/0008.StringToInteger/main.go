package leetcode

func myAtoi(s string) int {
	isBegin := true
	maxInt32 := int64(1<<31 - 1)
	minInt32 := int64(-(1 << 31))
	var flag int64 = 1
	var result int64
	for i, _ := range s {
		c := s[i]
		if isBegin {
			if c == ' ' {
				continue
			} else if c == '+' {
				flag = 1
				isBegin = false
			} else if c == '-' {
				flag = -1
				isBegin = false
			} else if c >= '0' && c <= '9' {
				result = int64(c - '0')
				isBegin = false
			} else {
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				result = result*10 + int64(c-'0')
			} else {
				break
			}
		}
		if result > maxInt32 {
			break
		}
	}
	result *= flag

	if result < minInt32 {
		result = minInt32
	}
	if result > maxInt32 {
		result = maxInt32
	}
	return int(result)
}
