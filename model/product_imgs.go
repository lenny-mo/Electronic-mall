package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductId uint   `gorm:"not null;comment:商品ID;"`
	ImgPath   string `gorm:"not null;comment:图片地址;type:varchar(255)"`
}
