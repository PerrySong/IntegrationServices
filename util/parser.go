package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

/**
Fetch info from rest api and convert it to a struct
*/
func FetchFromWeb(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	fmt.Println("url = ", url)
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(url)
	fmt.Println("fetching from web = ", r)
	if r.StatusCode != 200 {
		err = errors.New("Fail to fetch user from github: " + r.Status)
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
