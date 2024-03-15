package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/auth"
	"take-a-break/web-service/constants"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthTokenHandler(t *testing.T) {

	test_token := constants.TEST_TOKEN

	// Create a new Gin context
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/auth/verify/:token", auth.GetAuthTokenHandler)

	// Providing Unauthorized JWT Token
	req, err := http.NewRequest(http.MethodGet, "/auth/verify/"+test_token, nil)
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
	if respBody["error"] != "Failed to parse the JWT." {
		t.Fatalf("Expected to get error '%s' but instead got '%s'\n", "Failed to parse the JWT.", respBody["error"])
	}
}
