package login

import (
	"net/http"
	"net/http/httptest"
	"take-a-break/web-service/login"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandleGoogleCallback(t *testing.T) {
	// Create a new Gin context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest("GET", "/GoogleCallback?state=random&code=validCode", nil)

	// Call the HandleGoogleCallback function
	login.HandleGoogleCallback(ctx)

	// For this part, I didn't figure out how to git a validcode in test, so I simply test if there will be a response.
	if ctx.Writer.Status() != http.StatusTemporaryRedirect {
		t.Errorf("Expected status code %d; got %d", http.StatusTemporaryRedirect, ctx.Writer.Status())
	}

}
