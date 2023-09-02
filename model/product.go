package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"not null;comment:商品名称;type:varchar(20)"`
	CatetoryID    uint   `gorm:"not null;comment:分类ID;"`
	Title         string `gorm:"not null;comment:商品标题;type:varchar(20)"` // 商品标题
	Info          string `gorm:"not null;comment:商品简介;type:varchar(255)"`
	ImgPath       string `gorm:"not null;comment:商品图片地址;type:varchar(255)"`
	Price         string `gorm:"not null;comment:商品价格;type:varchar(20)"`
	DiscountPrice string `gorm:"not null;comment:商品折扣价格;type:varchar(20)"`
	OnSale        bool   `gorm:"not null;comment:是否上架;default:false"`
	Num           uint   `gorm:"not null;comment:商品数量;default:0"`
	BossId        uint   `gorm:"not null;comment:商家ID;"`
	BossName      string `gorm:"not null;comment:商家名称;type:varchar(20)"`
	BossAvatar    string `gorm:"not null;comment:商家头像;type:varchar(255)"`
}

