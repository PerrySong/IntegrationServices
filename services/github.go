package services

import (
	"github.com/PerrySong/OAuth2/models"
)

func GetGithubToken(id uint64) (string, error) {
	token, err := models.GetGitToken(id)
	return token, err
}

//func GithubWholeInfo(id uint64) {
//	client := graphql.NewClient(githubGraphQlEndPoint)
//	req := graphql.NewRequest(`
//    query ($key: String!) {
//        items (id:$key) {
//            field1
//            field2
//            field3
//        }
//    }`)
//}
