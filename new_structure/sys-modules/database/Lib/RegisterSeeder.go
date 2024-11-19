package Lib

import (
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
)

type DbSeederInterface interface {
	Name() string
	Handle(db *gorm.DB)
}

var registeredSeeders = map[string]DbSeederInterface{}

func RegisterSeeder(seeder DbSeederInterface) {
	registeredSeeders[seeder.Name()] = seeder
}

func GetSeeders() map[string]DbSeederInterface {
	return registeredSeeders
}

func RegisterSeeders(seeders []DbSeederInterface) {
	for _, seeder := range seeders {
		RegisterSeeder(seeder)
	}
}

func GetSortedSeeders() []DbSeederInterface {
	seeders := make([]DbSeederInterface, 0, len(registeredSeeders))
	for _, seeder := range registeredSeeders {
		seeders = append(seeders, seeder)
	}

	sort.Slice(seeders, func(i, j int) bool {
		return extractNumber(seeders[i].Name()) < extractNumber(seeders[j].Name())
	})

	return seeders
}

func extractNumber(name string) int {
	parts := strings.SplitN(name, "_", 2)
	if len(parts) > 0 {
		num, err := strconv.Atoi(parts[0])
		if err == nil {
			return num
		}
	}
	return 0
}
