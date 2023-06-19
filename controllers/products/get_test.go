package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
)

func TestGetProduct(t *testing.T) {
	tests := []struct {
		w              *httptest.ResponseRecorder
		url            string
		expectedStatus int
		p              *models.Product
	}{
		{
			httptest.NewRecorder(),
			"/products/pid_test?locale=ar_ae",
			http.StatusOK,
			&models.Product{
				ID:          "pid_test",
				Name:        "Product Test",
				Description: "Product Test Description",
			},
		},
		{
			httptest.NewRecorder(),
			"/products/pid_test2?locale=ar_ae",
			http.StatusNotFound,
			&models.Product{},
		},
	}

	db.Init(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "at-cs", "changeme", "localhost", "at-cs_test"))
	cache.Init("localhost", "6379", "", 10)

	r := gin.Default()
	r.GET("/products/:id", GetProduct)

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, test.url, nil)

		if test.p != nil {
			db.DB.Create(test.p)
		}
		r.ServeHTTP(test.w, req)

		if test.w.Code != test.expectedStatus {
			t.Errorf("expected status code %d, got %d", test.expectedStatus, test.w.Code)
		}

		if test.w.Code == http.StatusOK {
			if test.w.Body.String() == "" {
				t.Errorf("expected body to be non-empty")
			}

			var p models.Product
			json.NewDecoder(test.w.Body).Decode(&p)

			if p.ID != test.p.ID {
				t.Errorf("expected product ID to be %s, got %s", test.p.ID, p.ID)
			}
		}
	}

}
