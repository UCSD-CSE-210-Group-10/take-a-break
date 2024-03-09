package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Config struct {
	ClientID        string
	ClientSecret    string
	AuthURL         string
	TokenURL        string
	RedirectURL     string
	ClientURL       string
	TokenSecret     string
	TokenExpiration int64
	PostURL         string
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

func getConfig() Config {
	return Config{
		ClientID:        os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:    os.Getenv("GOOGLE_CLIENT_SECRET"),
		AuthURL:         "https://accounts.google.com/o/oauth2/v2/auth",
		TokenURL:        "https://oauth2.googleapis.com/token",
		RedirectURL:     os.Getenv("REDIRECT_URL"),
		ClientURL:       os.Getenv("CLIENT_URL"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenExpiration: 36000,
		PostURL:         "https://jsonplaceholder.typicode.com/posts",
	}
}

func GetTokenParams(config Config, code string) string {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("client_secret", config.ClientSecret)
	params.Set("code", code)
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", config.RedirectURL)
	return params.Encode()
}

func isUCSDEmail(email string) bool {
	// Check if the email has the "ucsd.edu" domain
	return strings.HasSuffix(email, "ucsd.edu")
}

func GetAuthTokenHandler(c *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		fmt.Print("Hello")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error loading .env file"})
		return
	}

	config := getConfig()
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code must be provided"})
		return
	}

	tokenParams := GetTokenParams(config, code)

	resp, err := http.Post(config.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(tokenParams))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code for token"})
		return
	}
	defer resp.Body.Close()
	var tokenResp struct {
		IDToken string `json:"id_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token response"})
		return
	}

	fmt.Print(tokenResp)
	if tokenResp.IDToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth error"})
		return
	}

	jwksURL := "https://www.googleapis.com/oauth2/v3/certs"

	k, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a keyfunc.Keyfunc from the server's URL."})
		return
	}

	parsed, err := jwt.Parse(tokenResp.IDToken, k.Keyfunc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the JWT."})
		return
	}

	claims, _ := parsed.Claims.(jwt.MapClaims)

	user_email, ok := claims["email"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email claim not found in token"})
		return
	}

	authorized := isUCSDEmail(user_email)
	c.JSON(http.StatusOK, gin.H{"token": tokenResp.IDToken, "authorized": authorized})
}
