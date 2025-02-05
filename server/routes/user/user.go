package user_create

import (
	"cow_backend/routes/auth"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (s *User) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/user")
	apiGroup.POST("/create", s.Create())
	apiGroup.GET("/verifyEmail", s.VerifyEmail())
	apiGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.FederalOff, auth.RegionalOff, auth.Admin))
	apiGroup.GET("/whoami", s.Whoami())
}
