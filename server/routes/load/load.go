package load

import "github.com/gin-gonic/gin"

type Load struct{}

func (s *Load) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/load")
	apiGroup.POST("/cow", s.Cow())
}
