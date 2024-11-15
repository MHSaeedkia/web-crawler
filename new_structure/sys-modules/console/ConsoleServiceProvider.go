package console

import (
	"project-root/app"
	commands "project-root/sys-modules/console/Commands"
	"project-root/sys-modules/console/Lib"
)

func (s *ConsoleServiceProvider) Register() {
	Lib.RegisterCommand(&commands.ListCommand{})
}

func (s *ConsoleServiceProvider) Boot() {

}

type ConsoleServiceProvider struct{}

var _ app.ServiceProviderInterface = &ConsoleServiceProvider{}
