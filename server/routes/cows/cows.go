package cows

import "github.com/gin-gonic/gin"

type Cows struct {
}

func (c *Cows) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/cows")
	apiGroup.GET("/get", c.Get())
	apiGroup.POST("/filter", c.Filter())
}
