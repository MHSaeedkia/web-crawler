package utils

import "strings"

// ExtractContractAndPlaceType extracts the contract and place type from a URL.
func ExtractContractAndPlaceType(url string) (contractType, placeType string) {
	parts := strings.Split(url, "/")
	if len(parts) >= 6 {
		types := strings.Split(parts[5], "-")
		if len(types) >= 2 {
			contractType = types[0]
			placeType = types[1]
		}
	}
	return
}
