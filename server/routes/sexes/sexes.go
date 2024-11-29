package sexes

import "github.com/gin-gonic/gin"

type Sexes struct {
}

func (s *Sexes) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/sexes")
	apiGroup.GET("/", s.Get())
}
