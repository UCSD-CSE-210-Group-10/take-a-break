package auth

import (
	"github.com/gin-gonic/gin"
)

func GetAuthTokenHandler(c *gin.Context) {

	token := c.Param("token")

	VerifyJWTTokenLogin(c, token)
}
