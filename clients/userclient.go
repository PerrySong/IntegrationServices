package main

import (
	"context"
	"fmt"
	"github.com/PerrySong/OAuth2/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I am a user client")
	// TODO remove With insecure
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("client could not connect: %v", err)
	}
	defer cc.Close()

	client := userpb.NewUserServiceClient(cc)
	getGithub(client)
}

func getGithub(client userpb.UserServiceClient) {
	req := &userpb.UserGithubBasicRequest{
		Id: 1,
	}
	res, err := client.GetGithubInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling user rpc: %v", err)

	}
	log.Printf("res = %v", res)
}
