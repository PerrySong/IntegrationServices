package main

import "github.com/PerrySong/OAuth2/models"

func main() {
	db, err := models.GetDb()
	if err != nil {
		panic(err.Error())
	}
	models.Migrate(db, &models.GithubToken{})
	models.Migrate(db, &models.LinkedinToken{})
	models.Migrate(db, &models.Repository{})
	models.Migrate(db, &models.User{})
}
