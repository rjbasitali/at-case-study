package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.rjbasitali.com/at-case-study/pkg/auth"
	"github.com/gin-gonic/gin"
)

const (
	jwt_secret = "some_random_jwt_signing_secret_Q5JLLDl06dlcMEEFBr4e1OBZS5"
)

func TestAuthMiddleware(t *testing.T) {
	validToken, _ := auth.NewToken(jwt_secret, time.Hour)
	tests := []struct {
		w              *httptest.ResponseRecorder
		url            string
		token          string
		expectedStatus int
	}{
		{
			httptest.NewRecorder(),
			"/ping",
			"some_invalid_token",
			http.StatusForbidden,
		},
		{
			httptest.NewRecorder(),
			"/ping",
			"invalid_token",
			http.StatusForbidden,
		},
		{
			httptest.NewRecorder(),
			"/ping",
			"lorem-ipsum",
			http.StatusForbidden,
		},
		{
			httptest.NewRecorder(),
			"/ping",
			validToken,
			http.StatusOK,
		},
	}

	r := gin.Default()
	r.Use(AuthMiddleware(jwt_secret))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, test.url, nil)
		req.Header.Set("Authorization", test.token)

		r.ServeHTTP(test.w, req)

		if test.w.Code != test.expectedStatus {
			t.Errorf("expected status code %d, got %d", test.expectedStatus, test.w.Code)
		}
	}

}
