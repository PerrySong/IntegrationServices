package main

import (
	"context"
	"fmt"
	"github.com/PerrySong/OAuth2/models"
	"github.com/PerrySong/OAuth2/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	userpb.UserServiceServer
}

func (*server) GetGithubInfo(ctx context.Context, req *userpb.UserGithubBasicRequest) (*userpb.UserGithubBasicResponse, error) {
	fmt.Printf("GetGithubInfo function was invoked with %v", req)
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

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
