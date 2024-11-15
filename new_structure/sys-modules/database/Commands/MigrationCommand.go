package commands

import (
	"fmt"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/database/Facades"
	SysDatabase "project-root/sys-modules/database/Lib"
	"sort"
)

type MigrationCommand struct{}

func (c *MigrationCommand) Signature() string {
	return "database:migrate"
}

func (c *MigrationCommand) Description() string {
	return "Migrate all tables - need 'up' or 'down' param - for example database:migrate up"
}

func (c *MigrationCommand) Handle(args []string) {
	actionName, err := getActionNameAndValidation(args)
	isVerbose := getVerb(args)
	if err {
		return
	}
	// sort name a-z
	migrations := SysDatabase.GetMigrations()
	keys := make([]string, len(migrations))
	i := 0
	for k := range migrations {
		keys[i] = k
		i++
	}
	if actionName == "up" {
		sort.StringsAreSorted(keys)
	} else {
		sort.Strings(keys)
	}

	// call
	print("start migrate...", isVerbose)
	for _, key := range keys {
		migration := migrations[key]
		print("Migrating: "+migration.Name(), isVerbose)
		if actionName == "up" {
			migrations[key].Up(Facades.Db())
		} else {
			migrations[key].Down(Facades.Db())
		}

		print("Migrated:  "+migration.Name(), isVerbose)
	}
	print("successful migrate", isVerbose)
}

func print(msg string, isVerbose bool) {
	if isVerbose {
		fmt.Println(msg)
	}
}

func getActionNameAndValidation(args []string) (string, bool) {
	if len(args) < 1 {
		fmt.Println("set param 'up' or 'down'")
		return "", true
	}
	arg := args[0]
	if arg != "up" && arg != "down" {
		fmt.Println("set param 'up' or 'down'")
		return "", true
	}
	return arg, false
}

func getVerb(args []string) bool {
	if len(args) >= 2 {
		if args[1] == "-v" {
			return true
		}
	}
	return false
}

var _ Lib.CommandInterface = &MigrationCommand{}
