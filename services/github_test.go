package services

import "testing"

func TestFetchGithubInfo(t *testing.T) {
	_, err := GetGithubToken(1)
	if err != nil {
		t.Error(err)
	}
}
