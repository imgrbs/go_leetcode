package leetcode

import (
	"leetcode/internal"
	"testing"
)

func TestTwoSumWith2Len(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	actual := leetcode.TwoSum(nums, target)

	if len(actual) != 2 {
		t.Errorf("expect=%v actual=%v too many indices.", 2, len(actual))
	}

	if actual[0] != 0 {
		t.Errorf("expect=%v actual=%v 1st indices failed.", 0, actual)
	}

	if actual[1] != 1 {
		t.Errorf("expect=%v actual=%v 2nd indices failed.", 1, actual)
	}
}

func TestTwoSumWith3Len(t *testing.T) {
	nums := []int{1, 3, 3}
	target := 6

	actual := leetcode.TwoSum(nums, target)

	if len(actual) != 2 {
		t.Errorf("expect=%v actual=%v too many indices.", 2, len(actual))
	}

	if actual[0] != 1 {
		t.Errorf("expect=%v actual=%v 1st indices failed.", 1, actual)
	}

	if actual[1] != 2 {
		t.Errorf("expect=%v actual=%v 2nd indices failed.", 2, actual)
	}
}
func TestTwoSumWith3LenButSwapPosition(t *testing.T) {
	nums := []int{3, 2, 3}
	target := 6

	actual := leetcode.TwoSum(nums, target)

	if len(actual) != 2 {
		t.Errorf("expect=%v actual=%v too many indices.", 2, len(actual))
	}

	if actual[0] != 0 {
		t.Errorf("expect=%v actual=%v 1st indices failed.", 0, actual)
	}

	if actual[1] != 2 {
		t.Errorf("expect=%v actual=%v 2nd indices failed.", 2, actual)
	}
}
