package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserID"`
	UserID    uint    `gorm:"not null;comment:用户ID;"`
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint    `gorm:"not null;comment:商品ID;"`
	Boss      User    `gorm:"foreignKey:BossID"`
	BossID    uint    `gorm:"not null;comment:商家ID;"`
}
