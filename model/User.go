package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name string `gorm:"type:varchar(20);not null"`
    Email string `gorm:"type:varchar(20);not null"`
    PasswordDigest string `gorm:"type:varchar(20);not null"`
    Status string 
    Money string
    
}
