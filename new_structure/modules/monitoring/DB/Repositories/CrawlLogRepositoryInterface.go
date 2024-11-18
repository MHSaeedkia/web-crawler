package Repositories

import (
	Models "project-root/modules/monitoring/DB/Models"
)

type CrawlLogRepositoryInterface interface {
	FindByID(id int) (*Models.CrawlLog, error)
	List() ([]Models.CrawlLog, error)
	Create(crawlLog *Models.CrawlLog) error
	Update(crawlLog *Models.CrawlLog) error
	Delete(id int) error
	Truncate() error
}
