package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

var endpotin = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "604178843113-7u6pfrtmi5lsu89tuv2dlbp73h2dn71f.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8000/GoogleCallback",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: endpotin,
}

const oauthStateString = "random"

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
// 	state := r.FormValue("state")
// 	if state != oauthStateString {
// 		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}
// 	fmt.Println(state)

// 	code := r.FormValue("code")
// 	fmt.Println(code)
// 	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
// 	fmt.Println(token)
// 	if err != nil {
// 		fmt.Println("Code exchange failed with '%s'\n", err)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	fmt.Fprintf(w, "Content: %s\n", contents)
// }

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Failed to get user info: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read user info: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Check if user email ends with "ucsd.edu"
	if !strings.Contains(string(contents), "@ucsd.edu") {
		fmt.Fprintf(w, "You don't have permission to log in with this account.")
		return
	}

	fmt.Fprintf(w, "Content: %s\n", contents)
}
