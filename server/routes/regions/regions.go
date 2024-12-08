package regions

import (
	"github.com/gin-gonic/gin"
)

type Regions struct {
}

func (a *Regions) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/regions")
	apiGroup.GET("/:id", a.GetByID())
	apiGroup.GET("/", a.GetByFilter())
	apiGroup.GET("/:id/getFarms", a.GetFarms())
	apiGroup.GET("/:id/news", a.News())
}
