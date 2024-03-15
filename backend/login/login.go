package login

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/users"

	"github.com/gin-gonic/gin"
)

func GetConfig() models.Config {
	return models.Config{
		ClientID:        "256148397214-ivsp4a1ro4posdhv4iiot28mb3gb5s7n.apps.googleusercontent.com",
		ClientSecret:    "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",
		AuthURL:         "https://accounts.google.com/o/oauth2/v2/auth",
		TokenURL:        "https://oauth2.googleapis.com/token",
		RedirectURL:     "http://localhost:3000/",
		ClientURL:       "http://localhost:3000",
		TokenSecret:     "123456",
		TokenExpiration: 36000,
		PostURL:         "https://jsonplaceholder.typicode.com/posts",
	}
}

func GetTokenParams(config models.Config, code string) string {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("client_secret", config.ClientSecret)
	params.Set("code", code)
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", config.RedirectURL)
	return params.Encode()
}

func GetLoginHandler(c *gin.Context, conn *database.DBConnection) {
	config := GetConfig()
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
		user := models.User{
			EmailID: claims["email"].(string),
			Name:    claims["name"].(string),
			Role:    "user",
			Avatar:  claims["picture"].(string),
		}

		users.InsertUserIntoDatabase(conn, user)

	}
}
