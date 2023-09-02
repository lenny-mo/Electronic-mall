package model

import "gorm.io/gorm"

// 轮播图 
type Carousel struct {
	gorm.Model
	ImgUrl    string `gorm:"not null;comment:图片地址;type:varchar(255)"`
	ProductID uint   `gorm:"not null;comment:商品ID;"`
}




