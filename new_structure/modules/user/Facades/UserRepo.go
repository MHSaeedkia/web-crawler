package Facades

import (
	"project-root/modules/user/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func UserRepo() Repositories.UserRepositoryInterface {
	return &Repositories.UserRepository{
		Db: Facades.Db(),
	}
}
