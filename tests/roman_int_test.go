package leetcode

import (
	"leetcode/internal"
	"testing"
)

func Test_GiveIII_Want3(t *testing.T) {
	input := "III"
	expect := 3

	result := leetcode.RomanToInt(input)

	if result != expect {
		t.Errorf("Failed, assert %s got %d", input, result)
	}
}

func Test_GiveMCMXCIV_Want1994(t *testing.T) {
	input := "MCMXCIV"
	expect := 1994

	result := leetcode.RomanToInt(input)

	if result != expect {
		t.Errorf("Failed, assert %s got %d", input, result)
	}
}
