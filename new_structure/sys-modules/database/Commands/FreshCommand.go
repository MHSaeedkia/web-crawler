package commands

import (
	"project-root/sys-modules/console/Lib"
)

type FreshCommand struct{}

func (c *FreshCommand) Signature() string {
	return "database:fresh"
}

func (c *FreshCommand) Description() string {
	return "Executes these commands migrate up, migrate down, seed in order - To rewrite the tables"
}

func (c *FreshCommand) Handle(args []string) {
	Lib.CallManualCommand("database:migrate", []string{"down", "-v"})
	Lib.CallManualCommand("database:migrate", []string{"up", "-v"})
	Lib.CallManualCommand("database:seed", []string{"-v"})
}

var _ Lib.CommandInterface = &SeedCommand{}
