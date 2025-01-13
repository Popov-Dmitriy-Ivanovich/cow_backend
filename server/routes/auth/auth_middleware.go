package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
		fmt.Println("token:", token)
		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !jwtToken.Valid {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		isAuthorized := false
		userRole := GetRole(claims.RoleId)
		if userRole == Admin {
			isAuthorized = true
		}
		for _, role := range requiredRole {
			if userRole == role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			c.JSON(http.StatusForbidden, "У вас нет доступа к этому ресурсу")
			c.Abort()
			return
		}

		var RegionId = strconv.FormatUint(uint64(claims.RegionId), 10)
		farm := *claims.FarmId
		var FarmId = strconv.FormatUint(uint64(farm), 10)
		var UserId = strconv.FormatUint(uint64(claims.UserId), 10)
		var DistId = strconv.FormatUint(uint64(claims.Distid), 10)
		c.Set("RoleId", claims.RoleId)
		c.Set("RegionId", RegionId)
		c.Set("FarmId", FarmId)
		c.Set("UserId", UserId)
		c.Set("DistId", DistId)
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
