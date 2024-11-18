package Facades

import (
	"project-root/modules/source-site/DB/Repositories"
	"project-root/sys-modules/database/Facades"
)

func SourceSiteRepo() Repositories.SourceSiteRepositoryInterface {
	return &Repositories.SourceSiteRepository{
		Db: Facades.Db(),
	}
}
