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

func (repo *CrawlLogRepository) GetCrawlLogsWithPagination(perPage, pageNum int) (*[]Models.CrawlLog, int, error) {
	var crawlLogs []Models.CrawlLog
	var totalRecords int64
	if err := repo.Db.Model(&Models.CrawlLog{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int((totalRecords + int64(perPage) - 1) / int64(perPage)) // round up
	if pageNum < 1 {
		pageNum = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	offset := (pageNum - 1) * perPage

	// --
	if err := repo.Db.Order("created_at DESC").Limit(perPage).Offset(offset).Find(&crawlLogs).Error; err != nil {
		return nil, 0, err
	}

	return &crawlLogs, totalPages, nil
}

var _ CrawlLogRepositoryInterface = &CrawlLogRepository{}
