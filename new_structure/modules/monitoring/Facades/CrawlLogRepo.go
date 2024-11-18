package Facades

import (
	"project-root/modules/monitoring/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func CrawlLogRepo() Repositories.CrawlLogRepositoryInterface {
	return &Repositories.CrawlLogRepository{
		Db: Facades.Db(),
	}
}
