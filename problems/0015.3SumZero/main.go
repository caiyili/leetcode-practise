package leetcode

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	uniqNums := make([]int, 0)
	for _, n := range nums {
		uniqNums = append(uniqNums, n)
	}
	nums = uniqNums

	var result [][]int
	type Pair struct {
		i, j int
	}

	flipMap := make(map[int][]*Pair, 0)
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			sum := nums[i] + nums[j]
			if _, ok := flipMap[sum]; !ok {
				flipMap[sum] = make([]*Pair, 0)
			}
			flipMap[sum] = append(flipMap[sum], &Pair{i, j})
		}
	}
	fmt.Println(flipMap)

	resultMap := make(map[string]bool, 0)
	for i := 0; i < lenNums; i++ {
		target := 0 - nums[i]
		if pairs, ok := flipMap[target]; ok {
			for _, pair := range pairs {
				if i != pair.i && i != pair.j {
					threeNum := []int{nums[i], nums[pair.i], nums[pair.j]}
					sort.Ints(threeNum)
					key := fmt.Sprintf("%d,%d,%d", threeNum[0], threeNum[1], threeNum[2])
					if _, ok := resultMap[key]; !ok {
						result = append(result, threeNum)
						resultMap[key] = true
					}
				}
			}
		}
	}

	return result
}
