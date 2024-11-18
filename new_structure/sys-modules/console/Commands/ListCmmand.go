package commands

import (
	"fmt"
	"project-root/sys-modules/console/Lib"
)

type ListCommand struct{}

func (c *ListCommand) Signature() string {
	return "command:list"
}

func (c *ListCommand) Description() string {
	return "Display all registered commands"
}

func (c *ListCommand) Handle(args []string) {
	fmt.Println()
	fmt.Println("------------- List of commands: -------------")
	for _, command := range Lib.GetCommands() {
		fmt.Println(command.Signature(), " ----------> ", command.Description())
	}
	fmt.Println("------------- ------------- -------------")
	fmt.Println()
}

var _ Lib.CommandInterface = &ListCommand{}
