package leetcode

import ()

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var openPairMap = map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	var closePairMap = map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	var stack = make([]byte, 0)
	var top = 0
	for i := range s {
		c := s[i]
		if _, ok := openPairMap[c]; ok {
			stack = append(stack, c)
			top++
		} else {
			if top <= 0 {
				return false
			}
			top--
			popChar := stack[top]
			stack = stack[:top]
			matchChar := closePairMap[c]
			if popChar != matchChar {
				return false
			}
		}
	}
	return top == 0
}
