package login

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var endpotin = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "604178843113-7u6pfrtmi5lsu89tuv2dlbp73h2dn71f.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/GoogleCallback",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: endpotin,
}

const oauthStateString = "random"

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println("Google OAuth URL:", url)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Failed to get user info: %s\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read user info: %s\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Check if user email ends with "ucsd.edu"
	if !strings.Contains(string(contents), "@ucsd.edu") {
		fmt.Printf("not end with ucsd.edu")
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000?authorized=false")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
