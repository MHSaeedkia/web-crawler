package commands

import (
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/database/Facades"
	SysDatabase "project-root/sys-modules/database/Lib"
)

type SeedCommand struct{}

func (c *SeedCommand) Signature() string {
	return "database:seed"
}

func (c *SeedCommand) Description() string {
	return "Seed all tables - You can send -v parameter to display details"
}

func (c *SeedCommand) Handle(args []string) {
	isVerbose := hasVerb(args)

	// sort name a-z
	seeders := SysDatabase.GetSortedSeeders()
	// call
	print("start seed...", isVerbose)
	for _, seeder := range seeders {
		print("Seeding: "+seeder.Name(), isVerbose)
		seeder.Handle(Facades.Db())
		print("Seeded:  "+seeder.Name(), isVerbose)
	}
	print("successful seeders", isVerbose)
}

func hasVerb(args []string) bool {
	if len(args) == 1 {
		if args[0] == "-v" {
			return true
		}
	}
	return false
}

var _ Lib.CommandInterface = &SeedCommand{}
