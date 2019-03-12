package models

import "github.com/jinzhu/gorm"

type Repository struct {
	gorm.Model
	Name string
}

type User struct {

}

