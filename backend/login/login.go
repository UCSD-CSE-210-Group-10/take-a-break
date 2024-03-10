package login

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/users"

	"github.com/gin-gonic/gin"
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
		AuthURL:         os.Getenv("AUTHURL"),
		TokenURL:        os.Getenv("TOKENURL"),
		RedirectURL:     os.Getenv("REDIRECT_URL"),
		ClientURL:       os.Getenv("CLIENT_URL"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenExpiration: 36000,
		PostURL:         os.Getenv("POSTURL"),
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

func GetLoginHandler(c *gin.Context, conn *database.DBConnection) {

	err := godotenv.Load()
	if err != nil {
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

	if tokenResp.IDToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth error"})
		return
	}

	auth.VerifyJWTTokenLogin(c, tokenResp.IDToken)

	statusCode := c.Writer.Status()
	if statusCode == http.StatusOK {
		claims := auth.ReturnJWTToken(tokenResp.IDToken)
		user := users.User{
			EmailID: claims["email"].(string),
			Name:    claims["name"].(string),
			Role:    "user",
		}

		users.InsertUserIntoDatabase(conn, user)

	}
}
