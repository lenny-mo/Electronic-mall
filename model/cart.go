package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model         // 可以直接继承gorm.Model中的字段,例如ID,时间字段等,
	UserID     uint    `gorm:"not null;comment:用户ID;"`
	ProductID  uint    `gorm:"not null;comment:商品ID;"`
	BossID     uint    `gorm:"not null;comment:商家ID;"`
	Num        uint    `gorm:"not null;comment:商品数量;"`
	Price      float64 `gorm:"not null;comment:商品价格;"`
	MaxNum     uint    `gorm:"not null;comment:商品库存;"`
	Checked   bool    `gorm:"not null;comment:是否选中;"`
	
}
