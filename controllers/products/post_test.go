package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		w              *httptest.ResponseRecorder
		url            string
		expectedStatus int
		expectedBody   string
		p              *models.Product
	}{
		{
			httptest.NewRecorder(),
			"/products",
			http.StatusCreated,
			"{\"message\":\"Product created successfully\"}",
			&models.Product{
				ID:          "pid_test",
				Name:        "Product Test",
				Description: "Product Test Description",
			},
		},
		{
			httptest.NewRecorder(),
			"/products",
			http.StatusBadRequest,
			"{\"error\":\"invalid request body\"}",
			nil,
		},
		{
			httptest.NewRecorder(),
			"/products",
			http.StatusBadRequest,
			"{\"error\":\"missing required fields\"}",
			&models.Product{
				ID:          "",
				Name:        "Product Test",
				Description: "Product Test Description",
			},
		},
	}

	db.Init(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "at-cs", "changeme", "localhost", "at-cs_test"))
	cache.Init("localhost", "6379", "", 10)

	r := gin.Default()
	r.POST("/products", CreateProduct)

	for _, test := range tests {
		var body []byte
		if test.p != nil {
			body, _ = json.Marshal(test.p)
		}
		req, _ := http.NewRequest(http.MethodPost, test.url, strings.NewReader(string(body)))

		// delete product if exists
		if test.p != nil {
			db.DB.Delete(&test.p, "id = ?", test.p.ID)
		}

		r.ServeHTTP(test.w, req)

		if test.w.Code != test.expectedStatus {
			t.Errorf("expected status code %d, got %d, response %s", test.expectedStatus, test.w.Code, test.w.Body.String())
		}

		if test.w.Body.String() != test.expectedBody {
			t.Errorf("expected body to be %s, got %s", test.expectedBody, test.w.Body.String())
		}
	}

}
