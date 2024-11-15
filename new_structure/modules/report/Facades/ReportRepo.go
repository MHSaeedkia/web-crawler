package Facades

import (
	"project-root/modules/report/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func ReportRepo() Repositories.ReportRepositoryInterface {
	return &Repositories.ReportRepository{
		Db: Facades.Db(),
	}
}
