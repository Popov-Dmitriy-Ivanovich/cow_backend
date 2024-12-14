package auth

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (a *Auth) Test() func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims := &JwtClaims{}
		token = strings.TrimPrefix(token, "Bearer ")
		jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		res := claims.Role
		c.JSON(200, res)
	}
}
