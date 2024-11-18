package Repositories

import "project-root/modules/report/DB/Models"

type ReportFilterRepositoryInterface interface {
	Create(reportFilter *Models.ReportFilter) (*Models.ReportFilter, error)
	Update(reportFilter *Models.ReportFilter) error
	Delete(id int) error
	FindByID(id int) (*Models.ReportFilter, error)
	FindByReportId(reportID int) (*Models.ReportFilter, error)
	FindAll() ([]*Models.ReportFilter, error)
}
