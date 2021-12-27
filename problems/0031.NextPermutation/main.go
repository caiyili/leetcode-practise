package leetcode

import (
	"sort"
)

func nextPermutation(nums []int) {
	var length = len(nums)
	//从后往前，找到第一个降序的数
	var pos = -1
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pos = i
			break
		}
	}
	if pos < 0 {
		sort.Ints(nums)
		return
	}
	var minLargeIdx = -1
	for i := pos + 1; i < length; i++ {
		if nums[i] > nums[pos] {
			if minLargeIdx == -1 || nums[i] < nums[minLargeIdx] {
				minLargeIdx = i
			}
		}
	}
	var minLargeNum = nums[minLargeIdx]
	nums[minLargeIdx] = nums[pos]
	nums[pos] = minLargeNum
	sort.Ints(nums[pos+1:])
}
