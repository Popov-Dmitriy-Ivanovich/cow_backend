package auth

import "github.com/gin-gonic/gin"

func (a *Auth) Test() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "ok")
	}
}
