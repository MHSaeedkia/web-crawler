package commands

import (
	"fmt"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/database/Facades"
)

type MyCommand struct{}

func (c *MyCommand) Signature() string {
	return "user:my-command"
}

func (c *MyCommand) Description() string {
	return "create seed user"
}

func (c *MyCommand) Handle(args []string) {
	db := Facades.Db()
	//db.
	fmt.Println("Executing MyCommand with Param1:", db)

}

var _ Lib.CommandInterface = &MyCommand{}
