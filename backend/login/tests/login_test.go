package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"take-a-break/web-service/database"
	"take-a-break/web-service/login"
	"take-a-break/web-service/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetConfig(t *testing.T) {

	config := login.GetConfig()

	expectedConfig := models.Config{
		ClientID:        os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:    os.Getenv("GOOGLE_CLIENT_SECRET"),
		AuthURL:         os.Getenv("AUTHURL"),
		TokenURL:        os.Getenv("TOKENURL"),
		RedirectURL:     os.Getenv("REDIRECT_URL"),
		ClientURL:       os.Getenv("CLIENT_URL"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenExpiration: 36000,
		PostURL:         os.Getenv("POSTURL"),
	}

	if config != expectedConfig {
		t.Errorf("GetConfig() returned unexpected result.\nGot: %+v\nExpected: %+v", config, expectedConfig)
	}
}

func TestLoginNoAuthCode(t *testing.T) {

	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// Create a new Gin context
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/auth/token", func(c *gin.Context) {
		login.GetLoginHandler(c, conn)
	})

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

	conn, err := database.NewDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Create a new Gin context
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/auth/token", func(c *gin.Context) {
		login.GetLoginHandler(c, conn)
	})

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
		t.Fatalf("Expected to get error '%s' but instead got '%s'\n", "Auth error", respBody["error"])
	}
}
