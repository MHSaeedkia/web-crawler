package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/source-site/DB/Models"
)

type SourceSiteRepository struct {
	Db *gorm.DB
}

func (repo *SourceSiteRepository) FindByID(id int) (*Models.SourceSite, error) {
	var sourceSite Models.SourceSite
	if err := repo.Db.First(&sourceSite, id).Error; err != nil {
		return nil, err
	}
	return &sourceSite, nil
}

func (repo *SourceSiteRepository) List() ([]Models.SourceSite, error) {
	var sourceSites []Models.SourceSite
	if err := repo.Db.Find(&sourceSites).Error; err != nil {
		return nil, err
	}
	return sourceSites, nil
}

func (repo *SourceSiteRepository) Create(sourceSite *Models.SourceSite) error {
	return repo.Db.Create(sourceSite).Error
}

func (repo *SourceSiteRepository) Update(sourceSite *Models.SourceSite) error {
	return repo.Db.Save(sourceSite).Error
}

func (repo *SourceSiteRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.SourceSite{}, id).Error
}

func (repo *SourceSiteRepository) Truncate() error {
	return repo.Db.Where("1 = 1").Delete(&Models.SourceSite{}).Error
}

var _ SourceSiteRepositoryInterface = &SourceSiteRepository{}
