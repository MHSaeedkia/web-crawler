package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/user/DB/Models"
)

type TelSessionRepository struct {
	Db *gorm.DB
}

func (repo *TelSessionRepository) FindByChatID(chatID int64) (*Models.TelSession, error) {
	var telSession Models.TelSession
	if err := repo.Db.Preload("LoggedUser").Where("chat_id = ?", chatID).First(&telSession).Error; err != nil {
		return nil, err
	}
	return &telSession, nil
}

func (repo *TelSessionRepository) Create(telSession *Models.TelSession) (*Models.TelSession, error) {
	if err := repo.Db.Create(telSession).Error; err != nil {
		return nil, err
	}
	return telSession, nil
}

func (repo *TelSessionRepository) UpdateByChatID(chatID int64, updatedSession *Models.TelSession) error {
	return repo.Db.Model(&Models.TelSession{}).
		Where("chat_id = ?", chatID).
		Updates(updatedSession).Error
}

var _ TelSessionRepositoryInterface = &TelSessionRepository{}
