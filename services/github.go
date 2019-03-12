package services

import (
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func HandleLoginGithub(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(randomState)
	//fmt.Println("hereherer:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallbackGithub(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//fmt.Println(r)
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {

		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://api.github.com/user?access_token="+token.AccessToken)
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