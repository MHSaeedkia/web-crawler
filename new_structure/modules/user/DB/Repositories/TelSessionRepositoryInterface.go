package Repositories

import "project-root/modules/user/DB/Models"

type TelSessionRepositoryInterface interface {
	FindByChatID(chatID int64) (*Models.TelSession, error)
	Create(telSession *Models.TelSession) (*Models.TelSession, error)
	UpdateByChatID(chatID int64, updatedSession *Models.TelSession) error
}
