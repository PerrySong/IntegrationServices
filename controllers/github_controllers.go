package controllers

import (
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/services"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func GithubFetchUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	token, err := services.GetGithubToken(1)
	if err != nil {
		panic(err)
	}
	html := fmt.Sprintf("<html><body><div>token = %s</div></body></html>", token)
	fmt.Fprint(w, html)
}

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling call back from github")
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := githubOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {

		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err = models.StoreToken(token.AccessToken, 1)
	if err != nil {
		fmt.Println("err storing token ", err)
	}

	tokenStr, err := models.GetGitToken(1)
	if err != nil {
		fmt.Println("fail to get token", token.AccessToken)
	}

	fmt.Printf("Get token from user1, token = %s", tokenStr)

	resp, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)
	//fmt.Println("here")
	http.Redirect(w, r, "localhost:9090", http.StatusTemporaryRedirect)
	return
	if err != nil {
		fmt.Printf("could not create token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Response: %s", content)

}
