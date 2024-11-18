package Repositories

import "project-root/modules/report/DB/Models"

type ReportRepositoryInterface interface {
	Create(report *Models.Report) (*Models.Report, error)
	Update(report *Models.Report) error
	Delete(id int) error
	SoftDelete(id int) error
	FindReport(reportId int) (*Models.Report, error)
	FindReportUserByTitle(userID int, title string) (*Models.Report, error)
	GetReportsByUserIdWithPagination(userID, perPage, pageNum int) (*[]Models.Report, int, error)
}
