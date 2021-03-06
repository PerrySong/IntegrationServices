package main

import (
	"fmt"
	"github.com/PerrySong/OAuth2/controllers"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/services"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Auth server is running in 8080")
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", services.HandleLoginGoogle)
	http.HandleFunc("/callback", services.HandleCallbackGoogle)

	http.HandleFunc("/login/github", controllers.GithubLoginHandler)
	http.HandleFunc("/callbackGit", controllers.GithubCallbackHandler)

	http.HandleFunc("/login/linkedin", services.HandleLoginLinkedin)
	http.HandleFunc("/callback/linkedin", services.HandleCallbackLinkedin)
	http.HandleFunc("/gettoken", controllers.GithubFetchUserInfoHandler)

	http.HandleFunc("/transfer", transferPage)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Println(err)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="login">Google login</a> 
				<a href="login/github?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNwZnNjbXpAZ21haWwuY29tIiwiaWQiOjEsImlzcyI6ImNvdXJzZSJ9.i4ymF6PF1cvVbA6QQiNkPk4HhJIHWsGgTr5FhUQhrf0">Github login</a>
				<a href="login/linkedin">Linkedin login</a>
				</body></html>`

	fmt.Fprint(w, html)
}

func transferPage(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body>
					<form action="localhost:3000">
    					<input type="submit" value="go to your main page" />
					</form>
				</body></html>`

	fmt.Fprint(w, html)
}

func initDb() {

	db, err := models.GetDb()
	if err != nil {
		panic(err.Error())
	}
	models.Migrate(db, &models.GithubToken{})
	models.Migrate(db, &models.LinkedinToken{})
	models.Migrate(db, &models.Repository{})
	models.Migrate(db, &models.GithubUser{})
}
