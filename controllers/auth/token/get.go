package token

import (
	"fmt"
	"net/http"

	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/auth"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetToken generates a new token.
// It returns the token in the response body and a status code `http.StatusCreated`.
// An error is returned if the token could not be generated.
func GetToken(c *gin.Context) {
	tokenString, err := auth.NewToken(cfg.JWT_SIGNING_SERCRET, cfg.JWT_EXPIRY_DURATION)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}

	t := models.Token{
		Token: tokenString,
	}

	c.JSON(http.StatusCreated, &t)
}
