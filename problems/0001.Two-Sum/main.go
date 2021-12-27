package leetcode

func findTwoSum(nums []int, target int) []int {
	flip := make(map[int]int, len(nums))
	for i, n := range nums {
		left := target - n
		if idx, ok := flip[left]; ok {
			return []int{idx, i}
		}
		flip[n] = i
	}
	return []int{0, 0}
}
