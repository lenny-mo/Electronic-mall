package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"not null;comment:公告内容;type:varchar(255)"`
}
