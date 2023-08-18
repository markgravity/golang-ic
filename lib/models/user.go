package models

import "gorm.io/gorm"

type User struct {
	Base              Base   `gorm:"embedded;"`
	Email             string `gorm:"unique"`
	EncryptedPassword string
}

func (u *User) Create(db *gorm.DB) (err error) {
	return db.Create(u).Error
}
