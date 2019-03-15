package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type GithubToken struct {
	gorm.Model
	UserId uint64 `gorm:"UNIQUE_INDEX"`
	Token  string
}

func StoreToken(tokenStr string, userid uint64) error {
	db, err := GetDb()
	defer db.Close()

	fmt.Println("Creating token")
	if err != nil {
		return err
	}
	token := GithubToken{UserId: userid, Token: tokenStr}
	db.Debug().Create(&token)
	return nil
}

func GetGitToken(userid uint64) (string, error) {
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
