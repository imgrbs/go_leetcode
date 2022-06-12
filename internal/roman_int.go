package leetcode

var romanValues = make(map[string]int)

func initRomanValues() {
	romanValues["M"] = 1000
	romanValues["CM"] = 900
	romanValues["D"] = 500
	romanValues["CD"] = 400
	romanValues["C"] = 100
	romanValues["XC"] = 90
	romanValues["L"] = 50
	romanValues["XL"] = 40
	romanValues["X"] = 10
	romanValues["IX"] = 9
	romanValues["V"] = 5
	romanValues["IV"] = 4
	romanValues["I"] = 1
}

func RomanToInt(s string) int {
	initRomanValues()

	result := 0

	lenChar := len(s)
	for i := 0; i < lenChar; i++ {
		var nextChar string
		if i+1 >= lenChar {
			nextChar = ""
		} else {
			nextChar = string(s[i+1])
		}

		romanKey := string(s[i]) + nextChar

		value := romanValues[romanKey]
		if value == 0 {
			result += romanValues[string(s[i])]
		} else {
			result += value
			i++
		}
	}

	return result
}
