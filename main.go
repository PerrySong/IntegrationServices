package main

import (
	"fmt"
	"github.com/PerrySong/OAuth/services"
	"net/http"
)


func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", services.HandleLoginGoogle)
	http.HandleFunc("/callback", services.HandleCallbackGoogle)

	http.HandleFunc("/loginGit", services.HandleLoginGithub)
	http.HandleFunc("/callbackGit", services.HandleCallbackGithub)


	http.HandleFunc("/login/linkedin", services.HandleLoginLinkedin)
	http.HandleFunc("/callback/linkedin", services.HandleCallbackLinkedin)

	http.ListenAndServe(":8080", nil)
}
func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="login">Google login</a> 
				<a href="loginGit">Github login</a>
				<a href="login/linkedin">Linkedin login</a>
				</body></html>`

	fmt.Fprint(w, html)
}









