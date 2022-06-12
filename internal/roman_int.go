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
		currentChar := string(s[i])
		nextChar := getNextChar(s, i, lenChar)

		value := romanValues[currentChar+nextChar]
		if isSubtractValue(value) {
			result += value
			i++
			continue
		}

		result += romanValues[currentChar]
	}

	return result
}

func getNextChar(s string, i int, lenChar int) string {
	if isOutOfLength(i, lenChar) {
		return ""
	}
	return string(s[i+1])
}

func isSubtractValue(value int) bool {
	return value != 0
}

func isOutOfLength(i int, lenChar int) bool {
	return i+1 >= lenChar
}
