package services

import (
	"fmt"
	"golang.org/x/oauth2"

	"io/ioutil"
	"net/http"
)

func HandleLoginLinkedin(w http.ResponseWriter, r *http.Request) {
	url := linkedinOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallbackLinkedin(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}


	fmt.Println(r)
	token, err := linkedinOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))


	if err != nil {

		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := &http.Client{
	}
	req, _ := http.NewRequest("GET", "https://api.linkedin.com/v2/me", nil)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("Token = " + token.AccessToken)
	req.Header.Set("Authorization", "Bearer " + token.AccessToken)
	req.Host = "api.linkedin.com"
	resp, err := client.Do(req)


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