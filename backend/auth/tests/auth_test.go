package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/auth"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIsUCSDEmail(t *testing.T) {

	if !auth.IsUCSDEmail("test-email@ucsd.edu") {
		t.Fatalf("Not a UCSD Email!")
	}
}

func TestReturnJWTToken(t *testing.T) {

	test_token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjA4YmY1YzM3NzJkZDRlN2E3MjdhMTAxYmY1MjBmNjU3NWNhYzMyNmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTgwNTU4MDg2MDY2NjUwNzI5MzAiLCJoZCI6InVjc2QuZWR1IiwiZW1haWwiOiJhYnVkaGlyYWphQHVjc2QuZWR1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJ6Q1BWU2RPYnI4NVczT0dRYjlNdDBRIiwibmFtZSI6IkFubW9sIEJ1ZGhpcmFqYSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NKZFVsVkYwMmZoOTBOby1CR3JydVJMOS1rRDFPejNCLTFtM3l0Q19vY1g9czk2LWMiLCJnaXZlbl9uYW1lIjoiQW5tb2wiLCJmYW1pbHlfbmFtZSI6IkJ1ZGhpcmFqYSIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNzEwMDQ3OTE3LCJleHAiOjE3MTAwNTE1MTd9.eBiNMvJNXwhpkf0qA1o6EbVxGX7rfu9AqhHTUFKYaDuU9mfmaMGOVUBFTZXIK6XXsXqV5GInViaF1VsWzmbVePTTwzJD0-u9rzT5XLdhChzLcTEH2JvzDwH1QA63S_fcnkZj3d6ewN7OQ2zNeKCAyFbOiaQAnTARp54R5Qp9x4v9ns04qU76a1WNa_a-NkROpRtLyVIZMLJbs3L0yEsj-Lct3zV2F5YWXqHeU-u3i5Iivnj4nAO3TDzXC_hrQXt8yijFIUZowuXwt1A-VKxsOkP4AvnpGyB5OopEByLFA1TuwxScJrD56FUZTNclSIhUWAA3hsNiajkZ8R9ypLe54Q"

	claims := auth.ReturnJWTToken(test_token)
	email := claims["email"].(string)
	if email != "abudhiraja@ucsd.edu" {
		t.Fatalf("Unable to Parse Token!")
	}
}

func TestVerifyJWTTokenLogin(t *testing.T) {

	test_token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjA4YmY1YzM3NzJkZDRlN2E3MjdhMTAxYmY1MjBmNjU3NWNhYzMyNmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyNTYxNDgzOTcyMTQtaXZzcDRhMXJvNHBvc2RodjRpaW90MjhtYjNnYjVzN24uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTgwNTU4MDg2MDY2NjUwNzI5MzAiLCJoZCI6InVjc2QuZWR1IiwiZW1haWwiOiJhYnVkaGlyYWphQHVjc2QuZWR1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJ6Q1BWU2RPYnI4NVczT0dRYjlNdDBRIiwibmFtZSI6IkFubW9sIEJ1ZGhpcmFqYSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NKZFVsVkYwMmZoOTBOby1CR3JydVJMOS1rRDFPejNCLTFtM3l0Q19vY1g9czk2LWMiLCJnaXZlbl9uYW1lIjoiQW5tb2wiLCJmYW1pbHlfbmFtZSI6IkJ1ZGhpcmFqYSIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNzEwMDQ3OTE3LCJleHAiOjE3MTAwNTE1MTd9.eBiNMvJNXwhpkf0qA1o6EbVxGX7rfu9AqhHTUFKYaDuU9mfmaMGOVUBFTZXIK6XXsXqV5GInViaF1VsWzmbVePTTwzJD0-u9rzT5XLdhChzLcTEH2JvzDwH1QA63S_fcnkZj3d6ewN7OQ2zNeKCAyFbOiaQAnTARp54R5Qp9x4v9ns04qU76a1WNa_a-NkROpRtLyVIZMLJbs3L0yEsj-Lct3zV2F5YWXqHeU-u3i5Iivnj4nAO3TDzXC_hrQXt8yijFIUZowuXwt1A-VKxsOkP4AvnpGyB5OopEByLFA1TuwxScJrD56FUZTNclSIhUWAA3hsNiajkZ8R9ypLe54Q"

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
