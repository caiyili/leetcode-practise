package leetcode

func maxAreaV1(height []int) int {
	var maxArea = 0
	var length = len(height)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			area := (j - i) * min(height[i], height[j])
			maxArea = max(area, maxArea)
		}
	}
	return maxArea
}

var min = func(a, b int) int {
	if a < b {
		return a
	}
	return b
}
var max = func(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var length = len(height)
	var start, end = 0, length - 1
	var maxArea = 0
	for start < end {
		area := (end - start) * min(height[start], height[end])
		maxArea = max(maxArea, area)
		if height[start] < height[end] {
			start++
			for start < end && height[start] < height[start-1] {
				start++
			}
		} else {
			end--
			for end > start && height[end] < height[end+1] {
				end--
			}
		}
	}
	return maxArea
}
