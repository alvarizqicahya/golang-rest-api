package middleware

import (
	"fmt"
	"golang-rest-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	SECRET_KEY     = "secret"
	SIGNING_METHOD = "HS256"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(SIGNING_METHOD) != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if token != nil && err == nil {
		claims := jwt.MapClaims{}
		_, _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})

	} else {
		c.JSON(http.StatusUnauthorized, helper.BuildErrorResponse("not authorized", err.Error()))
		c.Abort()
	}

}
