package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := strconv.Itoa(x)
	for i := range nums {
		size := len(nums)
		if nums[i] != nums[size-1-i] {
			return false
		}
	}
  
	return true
}
