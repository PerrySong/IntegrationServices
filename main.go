package main

import (
	"fmt"
	"github.com/PerrySong/OAuth2/controllers"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/services"
	"net/http"
)

func main() {

	initDb()

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", services.HandleLoginGoogle)
	http.HandleFunc("/callback", services.HandleCallbackGoogle)

	http.HandleFunc("/loginGit", controllers.GithubLoginHandler)
	http.HandleFunc("/callbackGit", controllers.GithubCallbackHandler)

	http.HandleFunc("/login/linkedin", services.HandleLoginLinkedin)
	http.HandleFunc("/callback/linkedin", services.HandleCallbackLinkedin)
	http.HandleFunc("/gettoken", controllers.GithubFetchUserInfoHandler)

	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="login">Google login</a> 
				<a href="loginGit">Github login</a>
				<a href="login/linkedin">Linkedin login</a>
				</body></html>`

	fmt.Fprint(w, html)
}

func initDb() {

	db, err := models.GetDb()
	db.SingularTable(true)
	if err != nil {
		panic(err.Error())
	}
	models.Migrate(db, &models.GithubToken{})
	models.Migrate(db, &models.LinkedinToken{})
	models.Migrate(db, &models.Repository{})
	models.Migrate(db, &models.User{})
}
