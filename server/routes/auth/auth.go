package auth

import "github.com/gin-gonic/gin"

type Auth struct {
}

func (s *Auth) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/auth")
	apiGroup.POST("/login", s.Login())
	apiGroup.POST("/register", s.Register())
	apiGroup.GET("/checkEmail", s.CheckEmail())
	testGroup := apiGroup.Group("/test")
	testGroup.Use(AuthMiddleware(Admin))
	testGroup.GET("/test", s.Test())

}
