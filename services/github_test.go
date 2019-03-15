package services

import (
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"testing"
)

func TestFetchGithubInfo(t *testing.T) {
	_, err := GetGithubToken(1)
	if err != nil {
		t.Error(err)
	}
}

func TestFetchUserFromGithub(t *testing.T) {
	var user models.User
	err := FetchUserFromGithub(1, &user)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("user = ", user)
}

func TestStoreUser(t *testing.T) {
	var user models.User
	if err := FetchUserFromGithub(1, &user); err != nil {
		t.Fatalf("Error when fetching user from github")
	}
	if err := models.StoreUser(&user); err != nil {
		t.Fatalf("Error when storing user to psql")
	}
	fmt.Println(user)
}
