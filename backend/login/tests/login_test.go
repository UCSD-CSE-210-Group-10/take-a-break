package login

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/login"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginNoAuthCode(t *testing.T) {
	// Create a new Gin context
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/auth/token", login.GetAuthTokenHandler)

	// jsonBody := []byte(`{"code": "temp-code"}`)
	// bodyReader := bytes.NewReader(jsonBody)/Users/anmolbudhiraja/Desktop/take-a-break/backend/.env

	// Not Providing Authorization Code
	req, err := http.NewRequest(http.MethodGet, "/auth/token?code=", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	var respBody map[string]interface{}
	json.NewDecoder(w.Body).Decode(&respBody)

	// Check to see if the response was what you expected
	if respBody["error"] != "Authorization code must be provided" {
		t.Fatalf("Expected to get error '%s' but instead got '%s'\n", "Authorization code must be provided", respBody["error"])
	}
}

func TestLoginInvalidAuthCode(t *testing.T) {
	// Create a new Gin context
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/auth/token", login.GetAuthTokenHandler)

	// jsonBody := []byte(`{"code": "temp-code"}`)
	// bodyReader := bytes.NewReader(jsonBody)/Users/anmolbudhiraja/Desktop/take-a-break/backend/.env

	// Not Providing Authorization Code
	req, err := http.NewRequest(http.MethodGet, "/auth/token?code=test-code", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	var respBody map[string]interface{}
	json.NewDecoder(w.Body).Decode(&respBody)

	// Check to see if the response was what you expected
	if respBody["error"] != "Auth error" {
		t.Fatalf("Expected to get error '%s' but instead got '%s'\n", "Authorization code must be provided", respBody["error"])
	}
}

// func TestIsUCSDEmail(t *testing.T) {

// 	if !login.isUCSDEmail("test-email@ucsd.edu") {
// 		t.Fatalf("")
// 	}
// }
