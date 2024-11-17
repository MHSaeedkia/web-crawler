package auth

import (
	"project-root/app"
	"project-root/modules/auth/Http/Controllers"
	"project-root/modules/auth/Http/Controllers/Login"
	"project-root/modules/auth/Http/Controllers/Register"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *AuthServiceProvider) Register() {
	Page.RegisterPages([]Page.PageInterface{
		// main
		&Controllers.MainAuthPage{},

		// register
		&Register.UsernameRegisterAuthPage{},
		&Register.PasswordRegisterAuthPage{},
		&Register.EmailRegisterAuthPage{},

		// login
		&Login.UsernameLoginAuthPage{},
		&Login.PasswordLoginAuthPage{},
	})
}

func (s *AuthServiceProvider) Boot() {

}

type AuthServiceProvider struct{}

var _ app.ServiceProviderInterface = &AuthServiceProvider{}
