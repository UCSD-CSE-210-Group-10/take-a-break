package users

import (
	"net/http"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/database"
	"take-a-break/web-service/models"
	"take-a-break/web-service/utils"

	"github.com/gin-gonic/gin"
)

type User = models.User

func PostUser(c *gin.Context, conn *database.DBConnection) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleBadRequest(c, "Failed to parse the request body", err)
		return
	}
	user, err := InsertUserIntoDatabase(conn, user)
	if err != nil {
		utils.HandleInternalServerError(c, "Failed to insert the user into the database", err)
		return
	}
	c.JSON(201, user)
}

func GetUserByEmailID(c *gin.Context, conn *database.DBConnection) (User, error) {

	token := c.Param("token")
	if !auth.VerifyJWTToken(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth Error"})
		return User{}, nil
	}

	claims := auth.ReturnJWTToken(token)

	emailID := claims["email"].(string)

	user, err := FetchUserByEmailID(conn, emailID)
	if err != nil {
		return User{}, err
	}
	c.JSON(200, user)
	return user, nil
}
