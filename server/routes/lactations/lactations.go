package lactations

import "github.com/gin-gonic/gin"

type Lactations struct {
}

func (l *Lactations) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/lactations")
	apiGroup.GET("/:id", l.Get())
}
