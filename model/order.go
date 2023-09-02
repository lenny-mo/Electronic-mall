package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNo   string  `gorm:"not null;comment:订单号;type:varchar(255)"`
	UserId    uint    `gorm:"not null;comment:用户ID;"`
	ProductId uint    `gorm:"not null;comment:商品ID;"`
	BossId    uint    `gorm:"not null;comment:商家ID;"`
	AddressId uint    `gorm:"not null;comment:地址ID;"`
	Num       uint    `gorm:"not null;comment:数量;"`
	OrderNum  uint    `gorm:"not null;comment:订单数量;"`
	Type      uint    `gorm:"not null;comment:订单类型;"` // 1:待付款 2:待发货 3:待收货 4:待评价 5:已完成 6:已取消
	Money     float64 `gorm:"not null;comment:订单金额;"`
}

