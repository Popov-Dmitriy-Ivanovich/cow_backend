package partners

import "github.com/gin-gonic/gin"

type Partners struct {
}

func (p *Partners) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/partners")
	apiGroup.GET("/", p.Get())
}
