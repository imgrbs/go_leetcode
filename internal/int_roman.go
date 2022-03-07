package leetcode

const MinimumValue = 0

var numberToRoman = make(map[int]string)
var romanKeys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func initRoman() {
	numberToRoman[1000] = "M"
	numberToRoman[900] = "CM"
	numberToRoman[500] = "D"
	numberToRoman[400] = "CD"
	numberToRoman[100] = "C"
	numberToRoman[90] = "XC"
	numberToRoman[50] = "L"
	numberToRoman[40] = "XL"
	numberToRoman[10] = "X"
	numberToRoman[9] = "IX"
	numberToRoman[5] = "V"
	numberToRoman[4] = "IV"
	numberToRoman[1] = "I"
}

func buildRoman(number *int, romanKey int, roman *string) {
	if isValidRomanKey(number, romanKey) {
		*number -= romanKey
		*roman += numberToRoman[romanKey]
	}

	if isSameRomanKey(number, romanKey) {
		buildRoman(number, romanKey, roman)
	}
}

func isSameRomanKey(number *int, romanKey int) bool {
	return *number >= romanKey
}

func isValidRomanKey(number *int, romanKey int) bool {
	return *number-romanKey >= MinimumValue
}

func IntToRoman(number int) string {
	initRoman()

	roman := ""

	for number > 0 {
		for _, romanKey := range romanKeys {
			buildRoman(&number, romanKey, &roman)
		}
	}

	return roman
}
