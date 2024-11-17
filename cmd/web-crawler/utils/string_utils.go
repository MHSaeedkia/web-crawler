package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
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


func ParseTimeFromString(input string) time.Time {
	input = strings.TrimSpace(input)

	now := time.Now()
	year, month, day := now.Date()

	createDate := func(y int, m time.Month, d int) time.Time {
		return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	}

	if strings.Contains(input, "لحظاتی پیش") || strings.Contains(input, "دقایقی پیش") {
		return createDate(year, month, day)
	}

	if strings.Contains(input, "دیروز") {
		return createDate(year, month, day-1)
	}

	if strings.Contains(input, "دو روز پیش") {
		return createDate(year, month, day-2)
	}

	if strings.Contains(input, "سه روز پیش") {
		return createDate(year, month, day-3)
	}

	if strings.Contains(input, "چهار روز پیش") {
		return createDate(year, month, day-4)
	}

	if strings.Contains(input, "پنج روز پیش") {
		return createDate(year, month, day-6)
	}
	if strings.Contains(input, "شش روز پیش") {
		return createDate(year, month, day-7)
	}
	if strings.Contains(input, "هفت روز پیش") {
		return createDate(year, month-1, day)
	}
	if strings.Contains(input, "یک ماه پیش") {
		return createDate(year, month-1, day)
	}

	if strings.Contains(input, "دو ماه پیش") {
		return createDate(year, month-2, day)
	}

	if strings.Contains(input, "سه ماه پیش") {
		return createDate(year, month-3, day)
	}

	return createDate(year, month-4, day)
}