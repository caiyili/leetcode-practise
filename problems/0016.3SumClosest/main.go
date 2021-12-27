package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var length = len(nums)
	type Result struct {
		sum  int
		diff int
		init bool
		nums []int
	}
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -1 * a
	}
	result := &Result{0, 0, false, []int{}}
	for index := 1; index < length-1; index++ {
		var start, end = 0, length - 1
		for start < index && index < end {
			sum := nums[start] + nums[index] + nums[end]
			//fmt.Printf("sum:%d nums[start]{%d} + nums[index]{%d} + nums[end]{%d}\n", sum, nums[start], nums[index], nums[end])
			diff := sum - target
			absDiff := abs(diff)
			if !result.init {
				result.init = true
				result.sum = sum
				result.diff = absDiff
				result.nums = []int{nums[start], nums[index], nums[end]}
			} else {
				if absDiff < result.diff {
					result.diff = absDiff
					result.sum = sum
					result.nums = []int{nums[start], nums[index], nums[end]}
				}
			}
			if diff > 0 {
				end--
				for end > index && nums[end] == nums[end+1] {
					end--
				}
			} else {
				start++
				for start < index && nums[start] == nums[start-1] {
					start++
				}
			}
		}
	}
	//fmt.Println(result.nums, result.sum)
	return result.sum
}
