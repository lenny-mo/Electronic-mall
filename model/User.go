package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"type:varchar(20);not null"`
	Email          string `gorm:"type:varchar(20);not null"`
	PasswordDigest string `gorm:"type:varchar(256);not null"`
	Status         string `gorm:"type:varchar(20);not null;default:'active'"`
	Money          string `gorm:"type:varchar(20);not null;default:'0'"`
	Avatar         string `gorm:"type:varchar(256);not null;default:'https://images3.alphacoders.com/132/1326573.jpeg'"`
}

const (
	PassWordCost = 12       //密码加密难度
	Active       = "active" //激活
)

// SetPassWord 设置密码
func (u *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}

	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
