package controllers

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/callback",
		// TODO: provide env variable
		ClientID:     "393312056306-032k0eoj5ecva38vjlrid0qc5igdqeut.apps.googleusercontent.com", //os.Getenv("GOOGLE_CLIENT_ID"), // environment variable
		ClientSecret: "bCNVTLUDxf0I7_q7I0YeiEOs",                                                 //os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	githubOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/callbackGit",
		// TODO: provide env variable
		ClientID:     "b6806ac237edfebb87ec",                     //os.Getenv("GOOGLE_CLIENT_ID"), // environment variable
		ClientSecret: "645b18a2062868fc17026f19b98f44961cdbc4ea", //os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"user:email", "repo", "gist"},
		Endpoint:     github.Endpoint,
	}

	linkedinOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback/linkedin",
		ClientID:     "86eknrzyll2ahw",
		ClientSecret: "u5oRBMf1M1mcG21g",
		Scopes:       []string{"r_liteprofile", "r_liteprofile", "w_share", "r_emailaddress", "rw_company_admin", "w_member_social"},
		Endpoint:     linkedin.Endpoint,
	}

	randomState           = "random"
	githubGraphQlEndPoint = "https://api.github.com/graphql"
)
