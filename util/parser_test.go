package util

import (
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"testing"
)

func TestFetchFromWeb(t *testing.T) {
	var user models.GithubUser
	if err := FetchFromWeb("https://api.github.com/user?access_token=41941f48f150dbc3ef8016bdf1403383da69ec5a", &user); err != nil {
		t.Error("Error wen fetch from web: ", err)
	}
	fmt.Println(user)
}
