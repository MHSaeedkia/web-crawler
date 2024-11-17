package Repositories

import (
	"project-root/modules/post/DB/Models"
	ReportModels "project-root/modules/report/DB/Models"
)

type PostRepositoryInterface interface {
	FindByID(id int) (*Models.Post, error)
	List() ([]Models.Post, error)
	Create(post *Models.Post) error
	Update(post *Models.Post) error
	Delete(id int) error
	Truncate() error
	FindBySourceSiteID(sourceSiteID int) ([]Models.Post, error)
	FindByUserID(userID int) ([]Models.Post, error)
	FindByStatus(status int) ([]Models.Post, error)
	GetPostsForFilter(filter *ReportModels.ReportFilter, perPage, pageNum int) (*[]Models.Post, int, error)
}
