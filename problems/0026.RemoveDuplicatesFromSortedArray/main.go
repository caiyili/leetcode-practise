package leetcode

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var last = nums[0]
	var count = 1
	for i := 1; i < length; i++ {
		if nums[i] != last {
			nums[count] = nums[i]
			count++
		}
		last = nums[i]
	}
	nums = nums[:count]
	fmt.Println(nums)
	return count
}
