package Repositories

import "project-root/modules/user/DB/Models"

type UserRepositoryInterface interface {
	FindByUsername(username string) (*Models.User, error)
	FindByID(id int) (*Models.User, error)
	List() ([]Models.User, error)
	Create(user *Models.User) error
	Update(user *Models.User) error
	Delete(id int) error
	Truncate() error
}
