package auth

import (
	"net/http"
	"strings"
	"take-a-break/web-service/constants"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsUCSDEmail(email string) bool {
	// Check if the email has the "ucsd.edu" domain
	return strings.HasSuffix(email, "ucsd.edu")
}

func VerifyJWTTokenLogin(c *gin.Context, token string) {
	jwksURL := constants.JWKSURL

	k, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a keyfunc.Keyfunc from the server's URL."})
		return
	}

	parsed, err := jwt.Parse(token, k.Keyfunc)
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

	authorized := IsUCSDEmail(user_email)

	c.JSON(http.StatusOK, gin.H{"token": token, "authorized": authorized})
}

func VerifyJWTToken(token string) bool {
	jwksURL := constants.JWKSURL

	k, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		return false
	}

	parsed, err := jwt.Parse(token, k.Keyfunc)
	if err != nil {
		return false
	}

	claims, _ := parsed.Claims.(jwt.MapClaims)

	user_email, ok := claims["email"].(string)
	if !ok {
		return false
	}

	authorized := IsUCSDEmail(user_email)
	return authorized
}

func ReturnJWTToken(token string) jwt.MapClaims {

	parsed, _ := jwt.Parse(token, nil)
	claims, _ := parsed.Claims.(jwt.MapClaims)

	return claims
}
