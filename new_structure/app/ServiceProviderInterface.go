package app

type ServiceProviderInterface interface {
	Register()
	Boot()
}
