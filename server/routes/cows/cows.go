package cows

import "github.com/gin-gonic/gin"

type Cows struct {
}

func (c *Cows) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/cows")
	apiGroup.GET("/:id", c.Get())
	apiGroup.GET("/:id/checkMilks", c.CheckMilks())
	apiGroup.POST("/filter", c.Filter())
}
