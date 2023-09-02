package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"not null;comment:分类名称;type:varchar(20)"`
}
