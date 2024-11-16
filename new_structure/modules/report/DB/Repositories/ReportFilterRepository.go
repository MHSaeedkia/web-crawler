package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
)

type ReportFilterRepository struct {
	Db *gorm.DB
}

func (repo *ReportFilterRepository) Create(reportFilter *Models.ReportFilter) (*Models.ReportFilter, error) {
	if err := repo.Db.Create(reportFilter).Error; err != nil {
		return nil, err
	}
	return reportFilter, nil
}

func (repo *ReportFilterRepository) Update(reportFilter *Models.ReportFilter) error {
	return repo.Db.Save(reportFilter).Error
}

func (repo *ReportFilterRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.ReportFilter{}, id).Error
}

func (repo *ReportFilterRepository) FindByID(id int) (*Models.ReportFilter, error) {
	var reportFilter Models.ReportFilter
	if err := repo.Db.First(&reportFilter, id).Error; err != nil {
		return nil, err
	}
	return &reportFilter, nil
}

func (repo *ReportFilterRepository) FindByReportId(reportID int) (*Models.ReportFilter, error) {
	var reportFilter Models.ReportFilter
	if err := repo.Db.Where("reports_id = ?", reportID).First(&reportFilter).Error; err != nil {
		return nil, err
	}
	return &reportFilter, nil
}

func (repo *ReportFilterRepository) FindAll() ([]*Models.ReportFilter, error) {
	var reportFilters []*Models.ReportFilter
	if err := repo.Db.Find(&reportFilters).Error; err != nil {
		return nil, err
	}
	return reportFilters, nil
}

var _ ReportFilterRepositoryInterface = &ReportFilterRepository{}
