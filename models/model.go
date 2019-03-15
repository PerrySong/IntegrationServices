package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate(db *gorm.DB, table interface{}) {
	if !db.HasTable(&table) {
		fmt.Println("Creating table")
		db.Debug().AutoMigrate(table)
	} else {
		db.AutoMigrate(&table)
	}
}

func GetDb() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	db.SingularTable(true)
	return db, err
}

func StoreUser(user *User) error {
	db, error := GetDb()
	if error != nil {
		return error
	}
	db.Debug().Create(&user)
	return nil
}
