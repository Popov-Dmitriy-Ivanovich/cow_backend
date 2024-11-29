package checkmilks

import "github.com/gin-gonic/gin"

type CheckMilks struct {
}

func (cm *CheckMilks) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/checkMilks")
	apiGroup.GET("/:id", cm.GetByID())
	apiGroup.GET("/", cm.GetByFilter())
}
