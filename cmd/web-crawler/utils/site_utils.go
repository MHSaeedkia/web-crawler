package utils

import (
	"log"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

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

// Helper function to get CPU usage
func GetCurrentCPUUsage() float64 {
	percentages, err := cpu.Percent(0, false)
	if err != nil || len(percentages) == 0 {
		log.Printf("Failed to get CPU usage: %v", err)
		return 0.0
	}
	return percentages[0] // Return CPU usage as a percentage
}

// Helper function to get RAM usage
func GetCurrentRAMUsage() int {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Failed to get RAM usage: %v", err)
		return 0
	}
	return int(vmStat.Used / (1024 * 1024)) // Convert bytes to MB
}
