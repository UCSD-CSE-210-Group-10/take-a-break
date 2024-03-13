package friends

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"take-a-break/web-service/constants"
// 	"take-a-break/web-service/database"
// 	"take-a-break/web-service/friends"
// 	"take-a-break/web-service/models"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSearchFriendsHandler(t *testing.T) {
// 	conn, err := database.NewDBConnection()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	router := gin.Default()
// 	test_token := constants.TEST_TOKEN

// 	router.GET("/friends/search/:token", friends.SearchFriendsHandler(conn))

// 	req, err := http.NewRequest("GET", "/friends/search/"+test_token+"?searchTerm=User", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code)

// 	var response []models.User
// 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.NotEmpty(t, response)
// }

// // Delete Friend is not called anywhere currently. Will add the test after using the functionality
// // func TestDeleteFriendHandler(t *testing.T) {

// // 	conn, err := database.NewDBConnection()
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	defer conn.Close()

// // 	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", "user3@example.com", "User 3", "user")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	_, err = conn.ExecuteQuery("INSERT INTO users (email_id, name, role) VALUES ($1, $2, $3)", "user4@example.com", "User 4", "user")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}

// // 	_, err = conn.ExecuteQuery("INSERT INTO friends (email_id1, email_id2) VALUES ($1, $2)", "user3@example.com", "user4@example.com")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}

// // 	router := gin.Default()

// // 	router.POST("/delete-friend", handle_friend.DeleteFriendHandler(conn))

// // 	payload := strings.NewReader("email_id1=user3@example.com&email_id2=user4@example.com")
// // 	req, err := http.NewRequest("POST", "/delete-friend", payload)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// // 	recorder := httptest.NewRecorder()

// // 	router.ServeHTTP(recorder, req)

// // 	assert.Equal(t, http.StatusOK, recorder.Code)

// // 	var response map[string]interface{}
// // 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}

// // 	assert.Contains(t, response["message"], "deleted successfully")
// // 	_, err = conn.ExecuteQuery("DELETE FROM friends WHERE email_id1 = $1 AND email_id2 = $2", "user3@example.com", "user4@example.com")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", "user3@example.com")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", "user4@example.com")
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // }
