package friend_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/constants"
	"take-a-break/web-service/database"
	"take-a-break/web-service/friend_request"
	"take-a-break/web-service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetFriendRequests(t *testing.T) {
	r := SetUpRouter()
	test_token := constants.TEST_TOKEN
	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// In testing mode
	r.GET("/friends/request/get/:token", func(c *gin.Context) {
		friend_request.GetFriendRequests(c, conn, true)
	})

	req, _ := http.NewRequest("GET", "/friends/request/get/"+test_token, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var requests []models.User
	json.Unmarshal(w.Body.Bytes(), &requests)
	var expectedRequests []models.User = []models.User{
		{
			EmailID: "user3@example.com",
			Name:    "Regular User 3",
			Role:    "user",
			Avatar:  "https://lh3.googleusercontent.com/a/ACg8ocJdUlVF02fh90No-BGrruRL9-kD1Oz3B-1m3ytC_ocX=s96-c",
		},
	}

	assert.Equal(t, requests, expectedRequests)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, requests)

	// Check Auth
	r = SetUpRouter()
	r.GET("/friends/request/get/:token", func(c *gin.Context) {
		friend_request.GetFriendRequests(c, conn)
	})

	req, _ = http.NewRequest("GET", "/friends/request/get/"+test_token, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var respBody map[string]interface{}
	json.NewDecoder(w.Body).Decode(&respBody)

	// Check to see if the response was what you expected
	if respBody["error"] != "Auth Error" {
		t.Fatalf("Expected to get error '%s' but instead got '%s'\n", "Auth Error", respBody["error"])
	}
}
