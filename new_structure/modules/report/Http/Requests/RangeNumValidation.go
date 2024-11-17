package Requests

import (
	"strconv"
	"strings"
)

func ParsAndValidateRangeNumbers(value string, min int, max int) (err bool, message string, num1 int, num2 int) {
	// Split the input based on ","
	parts := strings.Split(value, ",")
	if len(parts) != 2 {
		return true, "Please enter exactly two numbers separated by a comma", 0, 0
	}

	// Parse the two parts as integers
	num1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

	if err1 != nil || err2 != nil {
		return true, "Both values must be valid integers", 0, 0
	}

	// Check if the numbers are within the valid range (1 to 10) and num1 <= num2
	if num1 < min || num1 > max || num2 < min || num2 > max {
		return true, "Both numbers must be between 1 and 10", 0, 0
	}
	if num1 > num2 {
		return true, "The first number must be less than or equal to the second number", 0, 0
	}

	// No error, return the numbers
	return false, "", num1, num2
}
