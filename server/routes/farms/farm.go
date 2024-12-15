package farms

import (
	"cow_backend/models"
	"cow_backend/routes/auth"

	"github.com/gin-gonic/gin"
)

type Farms struct {
}

func (f *Farms) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/farms")
	authGroup := apiGroup.Group("")
	authGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	apiGroup.GET("/:id", f.GetByID())
	authGroup.GET("/", f.GetByFilter())
	authGroup.GET("/hoz", func(c *gin.Context) {

		db := models.GetDb()
		farms := []models.Farm{}
		qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id)").Find(&farms)
		if qres.Error != nil {
			c.JSON(500, qres.Error)
		}
		c.JSON(200, farms)
	})
}
