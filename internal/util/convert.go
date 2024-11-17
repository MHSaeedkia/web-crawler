package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var persianToEnglish = map[rune]rune{
	'۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4',
	'۵': '5', '۶': '6', '۷': '7', '۸': '8', '۹': '9',
}

func ConvertToInt(input string) int {
	var englishStr strings.Builder
	for _, char := range input {
		if englishDigit, ok := persianToEnglish[char]; ok {
			englishStr.WriteRune(englishDigit)
		} else if char >= '0' && char <= '9' {
			englishStr.WriteRune(char)
		}
	}
	reg, _ := regexp.Compile("[^0-9]+")
	englishStrStripped := reg.ReplaceAllString(englishStr.String(), "")

	num, err := strconv.Atoi(englishStrStripped)
	if err != nil {
		log.Printf("Error converting %s to int: %v", input, err)
		return 0
	}
	return num
}
