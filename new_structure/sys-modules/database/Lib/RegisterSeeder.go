package Lib

import "gorm.io/gorm"

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
