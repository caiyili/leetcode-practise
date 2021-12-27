package leetcode

import (
	"fmt"
	"sort"
)

func threeSumV2(nums []int) [][]int {
	sort.Ints(nums)

	var length = len(nums)
	var result = make([][]int, 0)
	var resultMap = make(map[string]bool, 0)
	for index := 1; index < length-1; index++ {
		start, end := 0, length-1
		for start < index && index < end {
			addSum := nums[index] + nums[start] + nums[end]
			if addSum == 0 {
				key := fmt.Sprintf("%d,%d,%d", nums[start], nums[index], nums[end])
				if _, ok := resultMap[key]; !ok {
					result = append(result, []int{nums[start], nums[index], nums[end]})
					resultMap[key] = true
				}
				for start+1 < index && nums[start] == nums[start+1] {
					start++
					continue
				}
				for end-1 >= index && nums[end] == nums[end-1] {
					end--
					continue
				}
				start++
				end--
			} else if addSum > 0 {
				end--
			} else if addSum < 0 {
				start++
			}
		}
	}

	return result
}
