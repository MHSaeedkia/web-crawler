package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// Persian to English digit map for converting numbers
var persianToEnglish = map[rune]rune{
	'۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4',
	'۵': '5', '۶': '6', '۷': '7', '۸': '8', '۹': '9',
}

// ConvertToInt converts a string with Persian or English digits to an integer.
func ConvertToInt(input string) int {
	var englishStr strings.Builder
	for _, char := range input {
		if englishDigit, ok := persianToEnglish[char]; ok {
			englishStr.WriteRune(englishDigit)
		} else if char >= '0' && char <= '9' {
			englishStr.WriteRune(char)
		}
	}

	// Remove any non-numeric characters
	reg, _ := regexp.Compile("[^0-9]+")
	englishStrStripped := reg.ReplaceAllString(englishStr.String(), "")
	number, err := strconv.Atoi(englishStrStripped)
	if err != nil {
		return 0
	}
	return number
}

// ConvertFloor extracts and converts the floor number from a Persian string.
func ConvertFloor(floor string) int {
	str := strings.Split(floor, " ")
	return ConvertToInt(str[0])
}


func ConvertFeatureToBool(feature string) bool {
	if strings.Contains(feature, "ندارد") {
		return false
	}
	return true
}