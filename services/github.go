package services

import (
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/util"
)

func GetGithubToken(id uint) (string, error) {
	token, err := models.GetGitToken(id)
	return token, err
}

func FetchReposFromGithub(userId uint, repos *[]models.GitRepository) error {
	token, err := GetGithubToken(userId)
	if err != nil {
		return err
	}
	err = util.FetchFromWeb("https://api.github.com/user/repos?access_token="+token, repos)
	return err
}

/*
	Fetch user from github and store it to db
*/
func FetchAndStoreRepos(userId uint) error {
	var repos []models.GitRepository
	err := FetchReposFromGithub(userId, &repos)
	if err != nil {
		return err
	}
	err = StoreRepos(userId, repos)
	if err != nil {
		return err
	}
	return nil
}

func FetchAndStoreUser(userId uint) error {
	var user models.GithubUser
	err := FetchUserFromGithub(userId, &user)
	if err != nil {
		return err
	}
	err = models.StoreUser(&user)
	return err
}

func FetchUserFromGithub(userId uint, user *models.GithubUser) error {
	token, err := GetGithubToken(userId)
	if err != nil {
		return err
	}
	err = util.FetchFromWeb("https://api.github.com/user?access_token="+token, user)
	user.UserId = uint(userId)
	return err
}

func StoreRepos(userId uint, repos []models.GitRepository) error {
	for i := 0; i < len(repos); i++ {
		repo := models.Repository{
			GitRepository: (repos)[i],
			UserId:        userId,
		}
		err := models.StoreRepo(&repo)
		if err != nil {
			return err
		}
	}
	return nil
}
