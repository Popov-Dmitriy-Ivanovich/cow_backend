package user_create

import "github.com/gin-gonic/gin"

type User struct {
}

func (s *User) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/user")
	apiGroup.POST("/create", s.Create())
	apiGroup.GET("/verifyEmail", s.VerifyEmail())
}
