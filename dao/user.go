package dao

import (
	"context"
	"eletronicMall/model"

	"gorm.io/gorm"
)

// 通常对应了数据库的一张表
type UserDAO struct {
	*gorm.DB
}

func NewUserDAO(ctx context.Context) *UserDAO {
	return &UserDAO{
		NewDBClient(ctx),
	}
}

func NewUserDAOWithDB(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db,
	}
}

// CheckUserExist 检查用户名是否存在
//
// 1. 如果用户名存在, 返回用户信息, true, nil
func (u *UserDAO) CheckUserExist(username string) (user *model.User, exist bool, err error) {

	var count int64
	err = u.DB.Model(&model.User{}).Where("name = ?", username).Find(&user).Count(&count).Error

	// 如果没有查询到用户，则返回nil, true, nil
	if count == 0 {
		return nil, false, err
	}

	return user, true, nil
}

// CreateUser 创建用户，使用事务
func (u *UserDAO) CreateUser(user *model.User) error {
	return u.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserByName 根据用户名获取用户信息
func (u *UserDAO) GetUserById(id uint) (user *model.User, err error) {
	err = u.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}


func (u *UserDAO) UpdateById(uid uint, user *model.User) error {
	return u.DB.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
}


 