package token

import (
	"fmt"
	"net/http"

	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/auth"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
)

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
