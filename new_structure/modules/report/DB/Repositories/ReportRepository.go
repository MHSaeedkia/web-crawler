package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
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

func (repo *ReportRepository) FindReportUserByTitle(userID int, title string) (*Models.Report, error) {
	var report Models.Report
	if err := repo.Db.Where("user_id = ? AND title = ?", userID, title).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

var _ ReportRepositoryInterface = &ReportRepository{}
