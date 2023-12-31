package products

import (
	"fmt"
	"net/http"

	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"git.rjbasitali.com/at-case-study/pkg/validate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProduct returns a product by id.
// It accepts a query parameter `locale` with value `en_ae` or `ar_ae` to return the product in the specified locale.
// It accepts a path parameter `id` to fetch the product.
// It returns the product in the response body and a status code `http.StatusOK`.
// An error is returned if the product could not be fetched or `locale` in invalid.
func GetProduct(c *gin.Context) {
	locale := c.Query("locale")
	if !validate.Locale(locale) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid locale",
		})
		return
	}

	id := c.Param("id")
	p := &models.Product{}

	if err := cache.Get(c, id, p); err == nil {
		c.JSON(http.StatusOK, p)
		return
	}

	result := db.DB.First(p, "id = ?", id)
	if result.Error != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error while fetching product", result.Error)
		switch result.Error {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	err := cache.Set(c, id, p)
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error while caching product", err)
	}

	c.JSON(http.StatusOK, p)
}
