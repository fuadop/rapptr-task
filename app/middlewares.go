package app

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (a *App) MustAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			a.HandleResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil, errors.New("unauthorized"))
			return
		}

		if strings.HasPrefix(authHeader, "Bearer") {
			tokens := strings.Split(authHeader, " ")
			authHeader = strings.Join(tokens[1:], " ")
		}

		authHeader = strings.TrimSpace(authHeader)

		claims := &JWTData{}
		_, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			a.HandleResponseJSON(c, http.StatusUnauthorized, "Invalid token", nil, err)
			return
		}

		user := claims.CustomClaims
		for k, v := range user {
			c.Set(k, v)
		}

		c.Next()
	}
}
