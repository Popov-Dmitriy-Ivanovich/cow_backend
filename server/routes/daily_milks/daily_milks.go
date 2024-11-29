package dailymilks

import "github.com/gin-gonic/gin"

type DailyMilk struct {
}

func (b *DailyMilk) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/dailyMilks")
	apiGroup.GET("/:id", b.Get())
}
