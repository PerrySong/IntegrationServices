package services

import (
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/util"
)

func GetGithubToken(id uint) (string, error) {
	token, err := models.GetGitToken(id)
	return token, err
}

func FetchUserFromGithub(id uint, user *models.User) error {
	token, err := GetGithubToken(id)
	if err != nil {
		return err
	}
	err = util.FetchFromWeb("https://api.github.com/user?access_token="+token, user)
	user.ID = uint(id)
	return err
}

func FetchReposFromGithub(id uint, repos *[]models.GitRepository) error {
	token, err := GetGithubToken(id)
	if err != nil {
		return err
	}
	err = util.FetchFromWeb("https://api.github.com/user/repos?access_token="+token, repos)
	return err
}
