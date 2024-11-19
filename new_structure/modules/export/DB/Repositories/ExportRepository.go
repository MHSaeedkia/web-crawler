package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/export/DB/Models"
)

type ExportRepository struct {
	Db *gorm.DB
}

func (repo *ExportRepository) FindByID(id int) (*Models.Export, error) {
	var export Models.Export
	if err := repo.Db.First(&export, id).Error; err != nil {
		return nil, err
	}
	return &export, nil
}

func (repo *ExportRepository) List() ([]Models.Export, error) {
	var exports []Models.Export
	if err := repo.Db.Find(&exports).Error; err != nil {
		return nil, err
	}
	return exports, nil
}

func (repo *ExportRepository) Create(export *Models.Export) error {
	return repo.Db.Create(export).Error
}

func (repo *ExportRepository) Update(export *Models.Export) error {
	return repo.Db.Save(export).Error
}

func (repo *ExportRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.Export{}, id).Error
}

func (repo *ExportRepository) Truncate() error {
	return repo.Db.Where("1 = 1").Delete(&Models.Export{}).Error
}

func (repo *ExportRepository) GetExportsWithPagination(perPage, pageNum int) (*[]Models.Export, int, error) {
	var exports []Models.Export
	var totalRecords int64
	if err := repo.Db.Model(&Models.Export{}).Count(&totalRecords).Error; err != nil {
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
	if err := repo.Db.Order("created_at DESC").Limit(perPage).Offset(offset).Find(&exports).Error; err != nil {
		return nil, 0, err
	}

	return &exports, totalPages, nil
}

var _ ExportRepositoryInterface = &ExportRepository{}
