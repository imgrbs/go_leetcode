package leetcode

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var x string
	currentX := l1
	for {
		if currentX.Next != nil {
			x = strconv.Itoa(currentX.Val) + x
			currentX = currentX.Next
		} else {
			break
		}
	}

	var y string
	currentY := l2
	for {
		if currentY.Next != nil {
			y = strconv.Itoa(currentY.Val) + y
			currentY = currentY.Next
		} else {
			break
		}
	}

	xVal, _ := strconv.Atoi(x)
	yVal, _ := strconv.Atoi(y)

	fmt.Println(xVal + yVal)

	return nil
}
