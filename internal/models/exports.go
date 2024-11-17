package models

import "time"

type Exports struct {
	ID        int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ReportID  int       `gorm:"column:report_id"`
	Report    Reports   `gorm:"foreignKey:ReportID;references:ID"`
	FileType  int       `gorm:"column:file_type;type:tinyint"`
	Status    int       `gorm:"column:status;type:tinyint"`
	FilePath  int       `gorm:"column:file_path;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
}
