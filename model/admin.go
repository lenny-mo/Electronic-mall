package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserID     uint   `gorm:"not null;comment:用户ID;"`
	PasswdHash string `gorm:"not null;comment:密码;type:varchar(255)"`
	Avatar     string `gorm:"not null;comment:头像;type:varchar(255)"`
}
