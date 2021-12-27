package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	type Window struct {
		left  int
		right int
		m     map[byte]int
	}
	findCharPos := func(w *Window, char byte) (int, bool) {
		if idx, ok := w.m[char]; ok {
			if idx >= w.left && idx <= w.right {
				return idx, true
			}
		}
		return 0, false
	}
	win := &Window{0, 0, make(map[byte]int, 1)}
	maxLength := 0
	for i, _ := range s {
		c := byte(s[i])
		if pos, ok := findCharPos(win, c); ok {
			win.left = pos + 1
		}
		win.right = i
		win.m[c] = i
		winLen := win.right - win.left + 1
		if winLen > maxLength {
			maxLength = winLen
		}
	}
	return maxLength
}
