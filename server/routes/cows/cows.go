package cows

import "github.com/gin-gonic/gin"

type Cows struct {
}

func (c *Cows) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/cows")
	apiGroup.GET("/", c.GetByFilter())
	apiGroup.GET("/:id", c.GetByID())
	apiGroup.GET("/:id/checkMilks", c.CheckMilks())
	apiGroup.GET("/:id/lactations", c.Lactations())
	apiGroup.GET("/:id/genetic", c.Genetic())
	apiGroup.GET("/:id/exterior", c.Exterior())
	apiGroup.GET("/:id/children", c.Children())
	apiGroup.GET("/:id/health", c.Health())
	apiGroup.POST("/filter", c.Filter())
}
