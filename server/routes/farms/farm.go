package farms

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

type Farms struct {
}

func (f *Farms) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/farms")
	apiGroup.GET("/:id", f.GetByID())
	apiGroup.GET("/", f.GetByFilter())
	apiGroup.GET("/hoz", func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id)").Find(&farms)
		if qres.Error != nil {
			c.JSON(500, qres.Error)
		}
		c.JSON(200, farms)
	})
}
