package gui

import "github.com/gin-gonic/gin"

type Gui struct {
}

func (s *Gui) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/gui")
	apiGroup.GET("/cowLoad", s.CowLoad())
	apiGroup.GET("", s.Index())
}
