package Facades

import (
	"project-root/modules/export/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func ExportRepo() Repositories.ExportRepositoryInterface {
	return &Repositories.ExportRepository{
		Db: Facades.Db(),
	}
}
