package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/constants"
	"take-a-break/web-service/database"
	"take-a-break/web-service/users"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// func TestPostUser(t *testing.T) {
// 	// Create a temporary database connection for testing
// 	conn, err := database.NewDBConnection()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Define the route
// 	router.POST("/users", func(c *gin.Context) {
// 		users.PostUser(c, conn)
// 	})

// 	// Create a sample user data for testing
// 	userData := users.User{
// 		EmailID: "testuser@example.com",
// 		Name:    "Test User",
// 		Role:    "user",
// 	}

// 	// Convert user data to JSON
// 	jsonData, err := json.Marshal(userData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a new HTTP request with the JSON data
// 	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a new HTTP recorder
// 	recorder := httptest.NewRecorder()

// 	// Handle the request
// 	router.ServeHTTP(recorder, req)

// 	// Check the response status code
// 	assert.Equal(t, http.StatusCreated, recorder.Code)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, "testuser@example.com", response["email_id"])
// 	assert.Equal(t, "Test User", response["name"])

// 	// Clean up the test data
// 	_, err = conn.ExecuteQuery("DELETE FROM users WHERE email_id = $1", userData.EmailID)
// 	assert.NoError(t, err, "Failed to clean up the test data")
// }

func TestGetUserByEmailID(t *testing.T) {

	test_token := constants.TEST_TOKEN
	conn, err := database.NewDBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define the route
	router.GET("/users/:token", func(c *gin.Context) {
		users.GetUserByEmailID(c, conn, true)
	})

	userData := users.User{
		EmailID: "abudhiraja@ucsd.edu",
		Name:    "Anmol Budhiraja",
		Role:    "user",
		Avatar:  "https://lh3.googleusercontent.com/a/ACg8ocJdUlVF02fh90No-BGrruRL9-kD1Oz3B-1m3ytC_ocX=s96-c",
	}

	// Create a new HTTP request to retrieve the user by email ID
	req := httptest.NewRequest("GET", "/users/"+test_token, nil)
	recorder := httptest.NewRecorder()

	// Handle the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response users.User
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// add check for user data and response
	assert.Equal(t, userData, response, "User data does not match")
}
