package analitics

import "github.com/gin-gonic/gin"

type CheckMilks struct{}

func (cm *CheckMilks) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/checkMilks")
	apiGroup.GET("/years", func(c *gin.Context) { c.JSON(200, "not done yet") })
}
