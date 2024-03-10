package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAuthTokenHandler(c *gin.Context) {

	token := c.Param("token")
	fmt.Print(token)

	VerifyJWTTokenLogin(c, token)
}
