package controllers

import (
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/services"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"log"
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

/*
	Request url:
	/login/github?jwt=[jwt]
*/
func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GithubLoginHandler")
	params := r.URL.Query()["jwt"]
	if len(params) == 0 {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	curState := params[0]
	url := githubOauthConfig.AuthCodeURL(curState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling call back from github")
	// getState

	// ... error handling

	// do something with decoded claims

	// TODO: Refactor check state
	// TODO: Origin check state:
	//if r.FormValue("state") != randomState {
	//	fmt.Println("state is not valid")
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}

	// Current check sate
	tokenString := r.FormValue("state")
	if len(tokenString) == 0 {
		log.Println("can not get state from github redirect")
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if claims["id"] == nil {
		log.Println("can not get id from github jwt mapclaims")
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	userId := uint(claims["id"].(float64))

	token, err := githubOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {

		log.Printf("can not get token from github code exchange %v \n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Store token
	err = models.StoreToken(token.AccessToken, userId)
	if err != nil {
		fmt.Println("err storing token ", err)
	}
	// Using stored token to fetch user

	services.FetchAndStoreRepos(uint(userId))
	services.FetchAndStoreUser(uint(userId))

	//resp, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)
	//log.Println("response = ", resp.Body)
	//
	//if err != nil {
	//	fmt.Printf("could not create token: %s\n", err.Error())
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}

	//defer resp.Body.Close()
	//content, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("could not parse response: %s\n", err.Error())
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}

	html := `<html><body>
				<a href="localhost:3000">Go to main page</a>
			</body></html>`

	fmt.Fprintf(w, html)

}
