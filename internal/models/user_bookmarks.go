package models

type UserBookMarks struct {
	UserID int    `gorm:"column:user_id"`
	User   Users  `gorm:"foreignKey:UserID;references:ID"`
	PostID int    `gorm:"column:post_id"` 
	Post   Posts  `gorm:"foreignKey:PostID;references:ID"`
}
