package Repositories

import "project-root/modules/source-site/DB/Models"

type SourceSiteRepositoryInterface interface {
	FindByID(id int) (*Models.SourceSite, error)
	List() ([]Models.SourceSite, error)
	Create(sourceSite *Models.SourceSite) error
	Update(sourceSite *Models.SourceSite) error
	Delete(id int) error
	Truncate() error
}
