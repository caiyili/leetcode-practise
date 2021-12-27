package leetcode

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	type Window struct {
		left, right int
	}
	getWindowLenth := func(w *Window) int {
		return w.right - w.left + 1
	}
	getWindowString := func(w *Window) string {
		return string(s[w.left : w.right+1])
	}
	getMaxWindow := func(a, b *Window) *Window {
		if getWindowLenth(a) > getWindowLenth(b) {
			return a
		}
		return b
	}

	lastWindow := make([]*Window, 0)
	maxWindow := &Window{0, 0}
	for i, _ := range s {
		arrWindow := make([]*Window, 0)
		win := &Window{i, i}
		arrWindow = append(arrWindow, win)
		if i > 0 && s[i] == s[i-1] {
			win := &Window{
				left:  i - 1,
				right: i,
			}
			arrWindow = append(arrWindow, win)
			maxWindow = getMaxWindow(win, maxWindow)
		}
		for _, win := range lastWindow {
			if win.left > 0 && s[win.left-1] == s[i] {
				win.left -= 1
				win.right = i
				arrWindow = append(arrWindow, win)
				maxWindow = getMaxWindow(win, maxWindow)
			}
		}
		lastWindow = arrWindow
	}
	return getWindowString(maxWindow)
}
