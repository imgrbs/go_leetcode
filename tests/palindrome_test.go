package leetcode

import (
	"testing"
)

func Test_10(t *testing.T) {
	result := IsPalindrome(10)

	if result == true {
		t.Errorf("Failed ")
	}
}

func Test_121(t *testing.T) {
	result := IsPalindrome(121)

	if result == false {
		t.Errorf("Failed ")
	}
}

func Test_11(t *testing.T) {
	result := IsPalindrome(11)

	if result == false {
		t.Errorf("Failed ")
	}
}

func Test_0(t *testing.T) {
	result := IsPalindrome(0)

	if result == false {
		t.Errorf("Failed ")
	}
}
