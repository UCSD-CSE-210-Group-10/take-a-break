package auth

import (
	"take-a-break/web-service/auth"
	"testing"
)

func TestIsUCSDEmail(t *testing.T) {

	if !auth.IsUCSDEmail("test-email@ucsd.edu") {
		t.Fatalf("Not a UCSD Email!")
	}
}
