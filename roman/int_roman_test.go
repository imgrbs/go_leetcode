package roman_test

import (
	roman2 "leetcode/roman"
	"testing"
)

func Test_Give1_WantI(t *testing.T) {
	input := 1
	expect := "I"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give5_WantV(t *testing.T) {
	input := 5
	expect := "V"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give4_WantIV(t *testing.T) {
	input := 4
	expect := "IV"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give9_WantIX(t *testing.T) {
	input := 9
	expect := "IX"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give3_WantIII(t *testing.T) {
	input := 3
	expect := "III"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give20_WantXX(t *testing.T) {
	input := 20
	expect := "XX"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give21_WantXXI(t *testing.T) {
	input := 21
	expect := "XXI"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give58_WantLVIII(t *testing.T) {
	input := 58
	expect := "LVIII"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}

func Test_Give1994_WantMCMXCIV(t *testing.T) {
	input := 1994
	expect := "MCMXCIV"

	result := roman2.IntToRoman(input)

	if result != expect {
		t.Errorf("Failed, assert %d got %s", input, result)
	}
}
