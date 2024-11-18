package Facades

import (
	"project-root/modules/post/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func PostRepo() Repositories.PostRepositoryInterface {
	return &Repositories.PostRepository{
		Db: Facades.Db(),
	}
}
