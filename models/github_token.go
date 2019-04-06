package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type GithubToken struct {
	gorm.Model
	UserId uint `gorm:"UNIQUE_INDEX"`
	Token  string
}

func StoreToken(tokenStr string, userid uint) error {
	db, err := GetDb()
	defer db.Close()

	fmt.Println("Creating token")
	if err != nil {
		return err
	}
	token := GithubToken{UserId: userid, Token: tokenStr}
	db.Debug().Model(&token).Update("token", token.Token)
	db.Debug().Create(&token)
	return nil
}

func GetGitToken(userid uint) (string, error) {
	db, err := GetDb()
	defer db.Close()
	if err != nil {
		return "", err
	}
	var token GithubToken
	db.Debug().Where("user_id = ?", userid).Find(&token)
	fmt.Println("token = " + token.Token)
	return token.Token, nil
}
