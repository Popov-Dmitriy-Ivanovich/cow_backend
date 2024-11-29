package districts

import "github.com/gin-gonic/gin"

type Districts struct {
}

func (b *Districts) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/districts")
	apiGroup.GET("/", b.Get())
}
