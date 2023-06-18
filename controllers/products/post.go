package products

import (
	"fmt"
	"net/http"

	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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
		fmt.Fprintln(gin.DefaultErrorWriter, "error while inserting new product", result.Error)
		if pgErr, ok := result.Error.(*pgconn.PgError); ok {
			switch pgErr.ConstraintName {
			case "products_pkey":
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

	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
