package regions

import (
	"github.com/gin-gonic/gin"
)

type Regions struct {
}

func (a *Regions) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/regions")
	apiGroup.GET("/get", a.Get())
}
