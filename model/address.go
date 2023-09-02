package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null;comment:用户ID;"`
	Name    string `gorm:"not null;comment:收货人姓名;type:varchar(20)"`
	Phone   string `gorm:"not null;comment:收货人手机号;type:varchar(11)"`
	Address string `gorm:"not null;comment:收货人地址;type:varchar(255)"`
}
