package Lib

type CommandInterface interface {
	Signature() string
	Description() string
	Handle(args []string)
}

var registeredCommands = make(map[string]CommandInterface)

func RegisterCommand(command CommandInterface) {
	registeredCommands[command.Signature()] = command
}

func GetCommands() map[string]CommandInterface {
	return registeredCommands
}

func GetCommand(name string) (CommandInterface, bool) {
	cmd, exists := registeredCommands[name]
	return cmd, exists
}

func CallManualCommand(name string, args []string) {
	cmd, exists := registeredCommands[name]
	if !exists {
		panic("Command not found")
	}
	cmd.Handle(args)
}
