package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/modules/user/Facades"
	"project-root/sys-modules/database/Lib"
	Lib2 "project-root/sys-modules/hash/Lib"
	"time"
)

type UserDbSeeder struct{}

func (s UserDbSeeder) Name() string {
	return "1_users_table"
}

func (s UserDbSeeder) Handle(db *gorm.DB) {
	Facades.UserRepo().Truncate()
	defaultPass, _ := Lib2.HashStr("123")

	Facades.UserRepo().Create(&Models.User{
		IsActive:      true,
		Username:      "admin",
		Password:      defaultPass,
		RoleType:      Enums.AdminRoleTypeEnum,
		Email:         "info@admin.com",
		CreatedChatID: nil,
		LastError:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		//LastError: time.Now(),
	})

	Facades.UserRepo().Create(&Models.User{
		IsActive:      true,
		Username:      "user",
		Password:      defaultPass,
		RoleType:      Enums.UserRoleTypeEnum,
		Email:         "info@user.com",
		CreatedChatID: nil,
		LastError:     time.Now(),
	})
}

var _ Lib.DbSeederInterface = &UserDbSeeder{}
