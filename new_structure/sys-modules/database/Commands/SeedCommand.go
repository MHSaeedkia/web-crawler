package commands

import (
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/database/Facades"
	SysDatabase "project-root/sys-modules/database/Lib"
	"sort"
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
	seeders := SysDatabase.GetSeeders()
	keys := make([]string, len(seeders))
	i := 0
	for k := range seeders {
		keys[i] = k
		i++
	}
	sort.StringsAreSorted(keys)

	// call
	print("start seed...", isVerbose)
	for _, key := range keys {
		seeder := seeders[key]
		print("Seeding: "+seeder.Name(), isVerbose)
		seeders[key].Handle(Facades.Db())
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
