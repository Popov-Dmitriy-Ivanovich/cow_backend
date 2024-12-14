package farms

import (
	"cow_backend/models"
	"cow_backend/routes/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Farms struct {
}

func (f *Farms) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/farms")
	apiGroup.GET("/:id", f.GetByID())
	apiGroup.GET("/", f.GetByFilter())
	authGroup := apiGroup.Group("")
	authGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	authGroup.GET("/hoz", func(c *gin.Context) {
		roleId, exists := c.Get("RoleId")
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}

		db := models.GetDb()
		farms := []models.Farm{}
		if roleId == 1 {
			farmId, exists := c.Get("FarmId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "FarmId не найден в контексте")
				return
			}
			qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id) AND id = ?", farmId).Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}

			c.JSON(200, farms)
		} else if roleId == 2 {
			regionId, exists := c.Get("RegionId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "RegionId не найден в контексте")
				return
			}
			qres := db.
				Joins("JOIN districts AS d ON farms.district_id = d.id").
				Joins("JOIN regions AS r ON r.id = d.region_id").
				Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id) AND r.id = ?", regionId).
				Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		} else {
			qres := db.Where("EXISTS (SELECT 1 FROM COWS WHERE cows.farm_group_id = farms.id)").Find(&farms)
			if qres.Error != nil {
				c.JSON(500, qres.Error)
			}
			c.JSON(200, farms)
		}

	})
}
