package friends

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/database"
	"take-a-break/web-service/friends"
	"take-a-break/web-service/models"
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
	test_token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjA4YmY1YzM3NzJkZDRlN2E3MjdhMTAxYmY1MjBmNjU3NWNhYzMyNmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTgwNTU4MDg2MDY2NjUwNzI5MzAiLCJoZCI6InVjc2QuZWR1IiwiZW1haWwiOiJhYnVkaGlyYWphQHVjc2QuZWR1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJ6Q1BWU2RPYnI4NVczT0dRYjlNdDBRIiwibmFtZSI6IkFubW9sIEJ1ZGhpcmFqYSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NKZFVsVkYwMmZoOTBOby1CR3JydVJMOS1rRDFPejNCLTFtM3l0Q19vY1g9czk2LWMiLCJnaXZlbl9uYW1lIjoiQW5tb2wiLCJmYW1pbHlfbmFtZSI6IkJ1ZGhpcmFqYSIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNzEwMDQ3OTE3LCJleHAiOjE3MTAwNTE1MTd9.eBiNMvJNXwhpkf0qA1o6EbVxGX7rfu9AqhHTUFKYaDuU9mfmaMGOVUBFTZXIK6XXsXqV5GInViaF1VsWzmbVePTTwzJD0-u9rzT5XLdhChzLcTEH2JvzDwH1QA63S_fcnkZj3d6ewN7OQ2zNeKCAyFbOiaQAnTARp54R5Qp9x4v9ns04qU76a1WNa_a-NkROpRtLyVIZMLJbs3L0yEsj-Lct3zV2F5YWXqHeU-u3i5Iivnj4nAO3TDzXC_hrQXt8yijFIUZowuXwt1A-VKxsOkP4AvnpGyB5OopEByLFA1TuwxScJrD56FUZTNclSIhUWAA3hsNiajkZ8R9ypLe54Q"

	router.GET("/friends/search/:token", friends.SearchFriendsHandler(conn))

	req, err := http.NewRequest("GET", "/friends/search/"+test_token+"?searchTerm=User", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, response)
}

// Delete Friend is not called anywhere currently. Will add the test after using the functionality
// func TestDeleteFriendHandler(t *testing.T) {

// 	conn, err := database.NewDBConnection()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", "user3@example.com", "User 3", "user")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", "user4@example.com", "User 4", "user")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_, err = conn.ExecuteQuery("INSERT INTO friends (email_id1, email_id2) VALUES ($1, $2)", "user3@example.com", "user4@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	router := gin.Default()

// 	router.POST("/delete-friend", handle_friend.DeleteFriendHandler(conn))

// 	payload := strings.NewReader("email_id1=user3@example.com&email_id2=user4@example.com")
// 	req, err := http.NewRequest("POST", "/delete-friend", payload)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Contains(t, response["message"], "deleted successfully")
// 	_, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 AND email_id2 = $2", "user3@example.com", "user4@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", "user3@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", "user4@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
