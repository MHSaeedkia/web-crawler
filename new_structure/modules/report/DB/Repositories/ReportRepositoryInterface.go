package Repositories

import "project-root/modules/report/DB/Models"

type ReportRepositoryInterface interface {
	Create(report *Models.Report) error
	Update(report *Models.Report) error
	Delete(id int) error
	FindReportUserByTitle(userID int, title string) (*Models.Report, error)
}
