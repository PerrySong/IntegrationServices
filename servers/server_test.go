package main

import (
	"context"
	"fmt"
	"github.com/PerrySong/OAuth2/proto"
	"testing"
)

func TestServer_GetGithubRepos(t *testing.T) {
	s := Server{}
	tests := []userpb.GithubRepoListRequest{
		{
			Id: 1,
		},
	}

	for _, req := range tests {

		resp, err := s.GetGithubRepos(context.Background(), &req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", err)
			return
		}
		if len(resp.Repos) == 0 {
			t.Errorf("Invalid repos %v", resp.Repos)
			return
		}

		fmt.Println("repos = ", resp.Repos)
	}
}

func TestServer_GetGithubInfo(t *testing.T) {
	s := Server{}
	tests := []userpb.UserGithubBasicRequest{
		{
			Id: 1,
		},
	}

	for _, req := range tests {

		resp, err := s.GetGithubInfo(context.Background(), &req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", err)
			return
		}
		if len(resp.Email) == 0 {
			t.Errorf("Invalid repos %v", resp.Email)
			return
		}

		fmt.Println("response = ", resp)
	}
}

func TestServer_UserHasToken(t *testing.T) {
	s := Server{}
	tests := []userpb.UserHasTokenRequest{
		{
			Id: 1,
		},
	}

	for _, req := range tests {

		resp, err := s.UserHasToken(context.Background(), &req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", err)
			return
		}
		if !resp.HasToken {
			t.Errorf("Invalid repos %v", resp)
			return
		}

		fmt.Println("response = ", resp)
	}

}
