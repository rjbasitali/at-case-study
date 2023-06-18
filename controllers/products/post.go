package products

import (
	"net/http"

	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/pkg/log"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	p := &models.Product{}
	if err := c.ShouldBindJSON(p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	result := db.DB.Create(p)
	if result.Error != nil {
		log.Error("error while inserting new product", result.Error)
		switch result.Error {
		case gorm.ErrDuplicatedKey:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "product already exists",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
