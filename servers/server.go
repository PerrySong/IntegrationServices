package main

import (
	"context"
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	userpb.UserServiceServer
}

func (*Server) GetGithubInfo(ctx context.Context, req *userpb.UserGithubBasicRequest) (*userpb.UserGithubBasicResponse, error) {
	id := req.GetId()
	db, err := models.GetDb()
	if err != nil {
		log.Fatalf("Can not get db %v", err)
	}
	var user models.GithubUser
	db.Debug().Where("user_id=?", id).First(&user)
	result := &userpb.UserGithubBasicResponse{
		Id:          int64(user.ID),
		Username:    user.UserName,
		Url:         user.Url,
		AvatarUrl:   user.AvatarUrl,
		Bio:         user.Bio,
		Company:     user.Company,
		ReposUrl:    user.ReposUrl,
		Email:       user.Email,
		Location:    user.Location,
		PublicRepos: int32(user.PublicRepos),
	}
	return result, nil
}

func (*Server) GetGithubRepos(ctx context.Context, req *userpb.GithubRepoListRequest) (*userpb.GithubRepoListResponse, error) {
	id := req.GetId()
	db, err := models.GetDb()
	if err != nil {
		log.Fatalf("Can not get db %v", err)
	}
	var repos []models.Repository
	db.Debug().Where("user_id=?", id).Find(&repos)
	var res userpb.GithubRepoListResponse
	for _, repo := range repos {
		repoProto := userpb.Repo{
			GitTagsUrl:      repo.GitTagsUrl,
			Description:     repo.Description,
			Private:         repo.Private,
			LanguagesUrl:    repo.LanguagesUrl,
			StargazersUrl:   repo.StargazersUrl,
			CommitsUrl:      repo.CommitsUrl,
			RepoCreatedAt:   &timestamp.Timestamp{Seconds: int64(repo.RepoCreatedAt.Second()), Nanos: int32(repo.RepoCreatedAt.Nanosecond())},
			RepoUpdatedAt:   &timestamp.Timestamp{Seconds: int64(repo.RepoUpdatedAt.Second()), Nanos: int32(repo.RepoUpdatedAt.Nanosecond())},
			HomePage:        repo.HomePage,
			StargazersCount: int32(repo.StargazersCount),
			LabelsUrl:       repo.LabelsUrl,
			Language:        repo.Language,
			Watchers:        int32(repo.Watchers),
		}
		res.Repos = append(res.Repos, &repoProto)
	}
	return &res, nil
}

func (*Server) UserHasToken(ctx context.Context, req *userpb.UserHasTokenRequest) (*userpb.UserHasTokenResponse, error) {
	id := req.GetId()
	var githubToken models.GithubToken
	db, err := models.GetDb()
	if err != nil {
		log.Fatalf("Can not get db %v", err)
	}
	db.Debug().Where("user_id=?", id).First(&githubToken)
	res := userpb.UserHasTokenResponse{HasToken: true}
	if len(githubToken.Token) == 0 {
		res.HasToken = false
		return &res, nil
	}
	return &res, nil
}

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
