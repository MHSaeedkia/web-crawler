package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/monitoring/DB/Models"
)

type CrawlLogRepository struct {
	Db *gorm.DB
}

func (repo *CrawlLogRepository) FindByID(id int) (*Models.CrawlLog, error) {
	var crawlLog Models.CrawlLog
	if err := repo.Db.First(&crawlLog, id).Error; err != nil {
		return nil, err
	}
	return &crawlLog, nil
}

func (repo *CrawlLogRepository) List() ([]Models.CrawlLog, error) {
	var crawlLogs []Models.CrawlLog
	if err := repo.Db.Find(&crawlLogs).Error; err != nil {
		return nil, err
	}
	return crawlLogs, nil
}

func (repo *CrawlLogRepository) Create(crawlLog *Models.CrawlLog) error {
	return repo.Db.Create(crawlLog).Error
}

func (repo *CrawlLogRepository) Update(crawlLog *Models.CrawlLog) error {
	return repo.Db.Save(crawlLog).Error
}

func (repo *CrawlLogRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.CrawlLog{}, id).Error
}

func (repo *CrawlLogRepository) Truncate() error {
	return repo.Db.Where("1 = 1").Delete(&Models.CrawlLog{}).Error
}

var _ CrawlLogRepositoryInterface = &CrawlLogRepository{}
