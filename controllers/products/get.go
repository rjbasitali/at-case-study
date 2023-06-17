package products

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.rjbasitali.com/at-case-study/pkg/models"
	"git.rjbasitali.com/at-case-study/pkg/validate"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	locale := c.Query("locale")
	if !validate.Locale(locale) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid locale",
		})
		return
	}

	id := c.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	product := models.Product{
		ID:          pid,
		Name:        fmt.Sprintf("Product %d", pid),
		Description: fmt.Sprintf("Product %d description", pid),
		CreatedAt:   time.Now(),
	}

	c.JSON(http.StatusOK, &product)
}
