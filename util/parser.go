package util

import (
	"encoding/json"
	"net/http"
	"time"
)

/**
Fetch info from rest api and convert it to a struct
*/
func FetchFromWeb(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
