package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
	"time"
)

type ReportRepository struct {
	Db *gorm.DB
}

func (repo *ReportRepository) Create(report *Models.Report) (*Models.Report, error) {
	if err := repo.Db.Create(report).Error; err != nil {
		return nil, err
	}
	return report, nil
}

func (repo *ReportRepository) Update(report *Models.Report) error {
	return repo.Db.Save(report).Error
}

func (repo *ReportRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.Report{}, id).Error
}

func (repo *ReportRepository) SoftDelete(id int) error {
	return repo.Db.Model(&Models.Report{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (repo *ReportRepository) FindReport(reportId int) (*Models.Report, error) {
	var report Models.Report
	if err := repo.Db.Where("id = ? AND deleted_at IS NULL", reportId).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (repo *ReportRepository) FindReportUserByTitle(userID int, title string) (*Models.Report, error) {
	var report Models.Report
	if err := repo.Db.Where("user_id = ? AND title = ?", userID, title).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (repo *ReportRepository) GetReportsByUserIdWithPagination(userID, perPage, pageNum int) (*[]Models.Report, int, error) {
	var reports []Models.Report
	var totalRecords int64

	if err := repo.Db.Model(&Models.Report{}).
		Where("users_id = ? AND deleted_at IS NULL", userID).
		Count(&totalRecords).Error; err != nil {
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

	if err := repo.Db.Where("users_id = ? AND deleted_at IS NULL", userID).
		Limit(perPage).
		Offset(offset).
		Find(&reports).Error; err != nil {
		return nil, 0, err
	}

	return &reports, totalPages, nil
}

var _ ReportRepositoryInterface = &ReportRepository{}
