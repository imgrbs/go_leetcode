package leetcode

func TwoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}
