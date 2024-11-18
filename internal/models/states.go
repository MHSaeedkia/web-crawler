package models

type States struct {
	ID            int    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Title         string `gorm:"column:title;type:varchar(255)"`
	TitleDivar    string `gorm:"column:title_divar;type:varchar(45)"`
	TitleSheypoor string `gorm:"column:title_sheypoor;type:varchar(45)"`
}
