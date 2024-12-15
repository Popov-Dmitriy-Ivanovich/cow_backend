package auth

import (
	"cow_backend/routes/admin"

	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func (s *Auth) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/auth")
	apiGroup.POST("/login", s.Login())
	apiGroup.GET("/checkEmail", s.CheckEmail())
	apiGroup.GET("/roles", s.Roles())
	adminGroup := apiGroup.Group("")
	adminGroup.Use(admin.AdminMiddleware())
	adminGroup.POST("/adminRegister", s.Register())
	testGroup := apiGroup.Group("/test")
	testGroup.Use(AuthMiddleware(Admin))
	testGroup.GET("/test", s.Test())

}
