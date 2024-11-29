package breeds

import "github.com/gin-gonic/gin"

type Breeds struct {
}

func (b *Breeds) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/breeds")
	apiGroup.GET("/:id", b.GetByID())
	apiGroup.GET("/", b.GetByFilter())
}
