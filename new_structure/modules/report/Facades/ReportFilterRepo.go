package Facades

import (
	"project-root/modules/report/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func ReportFilterRepo() Repositories.ReportFilterRepositoryInterface {
	return &Repositories.ReportFilterRepository{
		Db: Facades.Db(),
	}
}
