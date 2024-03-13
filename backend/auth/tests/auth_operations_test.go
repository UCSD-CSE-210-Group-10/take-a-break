package auth

import (
	"take-a-break/web-service/auth"
	"take-a-break/web-service/constants"
	"testing"
)

func TestIsUCSDEmail(t *testing.T) {

	if !auth.IsUCSDEmail("test-email@ucsd.edu") {
		t.Fatalf("Not a UCSD Email!")
	}
}

func TestReturnJWTToken(t *testing.T) {

	test_token := constants.TEST_TOKEN
	claims := auth.ReturnJWTToken(test_token)
	email := claims["email"].(string)
	if email != "abudhiraja@ucsd.edu" {
		t.Fatalf("Unable to Parse Token!")
	}
}

func TestVerifyJWTToken(t *testing.T) {

	test_token := constants.TEST_TOKEN
	authorized := auth.VerifyJWTToken(test_token)
	if authorized != false {
		t.Fatalf("Invalid Authentication!")
	}
}
