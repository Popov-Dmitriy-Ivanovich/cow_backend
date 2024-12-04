package auth

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(requiredRole []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims := &JwtClaims{}

		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil || !jwtToken.Valid {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		found := false
		for _, role := range requiredRole {
			if claims.Role == role {
				found = true
				break
			}
		}
		if found {
			c.Next()
			return
		} else {
			c.JSON(401, gin.H{"error": "no permission"})
			c.Abort()
			return
		}
	}
}
