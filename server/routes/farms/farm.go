package farms

import "github.com/gin-gonic/gin"

type Farms struct {
}

func (f *Farms) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/farms")
	apiGroup.GET("/:id", f.Get())
}
