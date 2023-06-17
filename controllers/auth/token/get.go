package token

import (
	"fmt"
	"net/http"
	"time"

	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(cfg.JWT_EXPIRY_DURATION).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(cfg.JWT_SIGNING_SERCRET))
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
