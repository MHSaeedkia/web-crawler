package Repositories

import (
	"project-root/modules/export/DB/Models"
)

type ExportRepositoryInterface interface {
	FindByID(id int) (*Models.Export, error)
	List() ([]Models.Export, error)
	Create(export *Models.Export) error
	Update(export *Models.Export) error
	Delete(id int) error
	Truncate() error
	GetExportsWithPagination(perPage, pageNum int) (*[]Models.Export, int, error)
}
