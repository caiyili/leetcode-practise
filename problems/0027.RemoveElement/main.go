package leetcode

func removeElement(nums []int, val int) int {
	length := len(nums)
	removeCount := 0

	for i := 0; i < length; i++ {
		if nums[i] == val {
			removeCount++
		} else {
			nums[i-removeCount] = nums[i]
		}
	}
	nums = nums[0 : length-removeCount]
	return length - removeCount
}
