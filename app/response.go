package app

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (a *App) HandleResponseJSON(c *gin.Context, code int, message string, data interface{}, err error) {
	res := gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	}
	if err != nil {
		res["error"] = err.Error()
	}

	c.JSON(code, res)
}

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom_claims"`
}

func (a *App) GenerateJWTAccessToken(userId, username string) (string, error) {
	// prepare claims for token
	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			// set token lifetime in timestamp
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		CustomClaims: map[string]string{
			"userId":   userId,
			"username": username,
		},
	}

	// generate a string using claims and HS256 algorithm
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the generated key using secretKey
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := tokenString.SignedString(secretKey)

	return token, err
}
