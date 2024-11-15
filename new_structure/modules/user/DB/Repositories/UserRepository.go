package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/user/DB/Models"
)

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) FindByUsername(username string) (*Models.User, error) {
	var user Models.User
	if err := repo.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByID(id int) (*Models.User, error) {
	var user Models.User
	if err := repo.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) List() ([]Models.User, error) {
	var users []Models.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(user *Models.User) error {
	return repo.Db.Create(user).Error
}

func (repo *UserRepository) Update(user *Models.User) error {
	return repo.Db.Save(user).Error
}

func (repo *UserRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.User{}, id).Error
}

func (repo *UserRepository) Truncate() error {
	return repo.Db.Where("1 = 1").Delete(&Models.User{}).Error
}
