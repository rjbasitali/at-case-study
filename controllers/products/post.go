package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
