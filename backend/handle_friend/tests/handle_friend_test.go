package handle_friend

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"take-a-break/web-service/database"
	"take-a-break/web-service/handle_friend"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchFriendsHandler(t *testing.T) {
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	router := gin.Default()

	router.GET("/search-friends", handle_friend.SearchFriendsHandler(conn))

	req, err := http.NewRequest("GET", "/search-friends?searchTerm=User", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []handle_friend.User
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response)
}

func TestDeleteFriendHandler(t *testing.T) {

	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	router := gin.Default()

	router.POST("/delete-friend", handle_friend.DeleteFriendHandler(conn))

	payload := strings.NewReader("email_id1=user1@example.com&email_id2=user2@example.com")
	req, err := http.NewRequest("POST", "/delete-friend", payload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, response["message"], "deleted successfully")
}
