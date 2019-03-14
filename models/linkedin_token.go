package models

import "github.com/jinzhu/gorm"

type LinkedinToken struct {
	gorm.Model
	UserId uint64 `gorm:"UNIQUE_INDEX"`
	Token string
}