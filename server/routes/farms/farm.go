package farms

import (
	"cow_backend/models"
	"cow_backend/routes/auth"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type Farms struct {
}

func (f *Farms) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/farms")
	authGroup := apiGroup.Group("")
	authGroup.GET("/hoz", func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}
		
		if err := db.Find(&farms, map[string]any{"type": 3}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, farms)
	})
	authGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	apiGroup.GET("/:id", f.GetByID())
	authGroup.GET("/", f.GetByFilter())
	
}
