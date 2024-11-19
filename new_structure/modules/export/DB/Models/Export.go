package Models

import (
	ReportModels "project-root/modules/report/DB/Models"
	"time"
)

type Export struct {
	ID        int                 `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ReportID  int                 `gorm:"column:report_id;type:mediumint"`
	Report    ReportModels.Report `gorm:"foreignKey:ReportID;references:ID"`
	FileType  int                 `gorm:"column:file_type;type:tinyint"`
	FilePath  string              `gorm:"column:file_path;type:text"`
	CreatedAt time.Time           `gorm:"column:created_at;type:datetime"`
}
