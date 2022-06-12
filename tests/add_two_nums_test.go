package leetcode

import (
	"leetcode/internal"
	"testing"
)

func Test243And564(t *testing.T) {
	l1v3 := leetcode.ListNode{3, nil}
	l1v2 := leetcode.ListNode{4, &l1v3}
	l1v1 := leetcode.ListNode{2, &l1v2}

	l2v3 := leetcode.ListNode{4, nil}
	l2v2 := leetcode.ListNode{6, &l2v3}
	l2v1 := leetcode.ListNode{5, &l2v2}

	expect3 := leetcode.ListNode{8, nil}
	expect2 := leetcode.ListNode{0, &expect3}
	expect1 := leetcode.ListNode{7, &expect2}

	result := leetcode.AddTwoNumbers(&l1v1, &l2v1)

	if result != &expect1 {
		t.Errorf("Failed, assert %d got %d", expect1.Val, result.Val)
	}
	if result.Next != &expect2 {
		t.Errorf("Failed, assert %d got %d", expect2.Val, result.Next.Val)
	}
	if result.Next.Next != &expect3 {
		t.Errorf("Failed, assert %d got %d", expect3.Val, result.Next.Next.Val)
	}
}
