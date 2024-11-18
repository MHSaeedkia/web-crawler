package Facades

import (
	"project-root/modules/user/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func TelSessionRepo() Repositories.TelSessionRepositoryInterface {
	return &Repositories.TelSessionRepository{
		Db: Facades.Db(),
	}
}
