package token

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
)

func TestGetToken(t *testing.T) {
	tests := []struct {
		w              *httptest.ResponseRecorder
		url            string
		expectedStatus int
	}{
		{
			httptest.NewRecorder(),
			"/auth/token",
			http.StatusCreated,
		},
		{
			httptest.NewRecorder(),
			"/invalid_url",
			http.StatusNotFound,
		},
	}

	r := gin.Default()
	r.GET("/auth/token", GetToken)

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, test.url, nil)

		r.ServeHTTP(test.w, req)

		if test.w.Code != test.expectedStatus {
			t.Errorf("expected status code %d, got %d", test.expectedStatus, test.w.Code)
		}

		if test.w.Code == http.StatusCreated {
			if test.w.Body.String() == "" {
				t.Errorf("expected body to be non-empty")
			}

			var token models.Token
			json.NewDecoder(test.w.Body).Decode(&token)

			if token.Token == "" {
				t.Errorf("expected token to be non-empty")
			}
		}
	}

}
