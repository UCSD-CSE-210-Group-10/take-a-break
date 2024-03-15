package login

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/users"

	"github.com/gin-gonic/gin"
)

func GetConfig() models.Config {
	return models.Config{
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
