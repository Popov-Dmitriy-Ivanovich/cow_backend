package monogenetic_illnesses

import "github.com/gin-gonic/gin"

type MonogeneticIllneses struct {
}

func (m *MonogeneticIllneses) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/monogenetic_illnesses")
	apiGroup.GET("/", m.Get())
}
