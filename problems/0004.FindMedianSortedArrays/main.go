package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	lenTotal := len1 + len2
	checkSep := func(midA int, midB int) int {
		if midA > 0 && midA < len1 && midB > 0 && midB < len2 {
			if nums1[midA-1] > nums2[midB] {
				return -1
			} else if nums1[midA] < nums2[midB-1] {
				return 1
			} else {
				return 0
			}
		}
		return 0
	}

	left, right := 0, len1
	midA := (right - left) / 2
	midB := lenTotal/2 - midA
	for {
		check := checkSep(midA, midB)
		if check < 0 {
			right = midA
			midA = (right - left) / 2
		} else if check > 0 {
			left = midA
			midA = (right - left) / 2
		} else {
			break
		}
		midB = lenTotal/2 - midA
	}

	if midA == 0 {
		//A里面所有数都大于B
	}

	return float64(0)
}
