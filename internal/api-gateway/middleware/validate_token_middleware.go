package middleware

import (
	"strings"

	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValidateAccessToken(accessSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientUnauthorized, "authorization header is required", nil, nil))
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(accessSecret), nil
		})

		if err != nil || !token.Valid {
			res := handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientUnauthorized, "invalid or expired token", nil, nil))
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
