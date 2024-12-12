package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type RoleType string

const (
	Farmer      RoleType = "farmer"
	RegionalOff RoleType = "regional"
	FederalOff  RoleType = "federal"
	Admin       RoleType = "admin"
)

func AuthMiddleware(requiredRole ...RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims := &JwtClaims{}
		token = strings.TrimPrefix(token, "Bearer ")
		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		fmt.Println(token)

		if err != nil || !jwtToken.Valid {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		isAuthorized := false
		userRole := GetRole(claims.Role)
		for _, role := range requiredRole {
			if userRole == role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет доступа к этому ресурсу"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetRole(roleId int) RoleType {
	switch roleId {
	case 1:
		return Farmer
	case 2:
		return RegionalOff
	case 3:
		return FederalOff
	case 4:
		return Admin
	default:
		return ""
	}
}
