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
	apiGroup.GET("/:id", f.GetByID())
	authGroup := apiGroup.Group("")
	authGroup.Use(auth.AuthMiddleware(auth.Farmer, auth.RegionalOff, auth.FederalOff))
	authGroup.GET("/", f.GetByFilter())
	apiGroup.GET("/hoz", func(c *gin.Context) {
		db := models.GetDb()
		farms := []models.Farm{}

		if err := db.Find(&farms, map[string]any{"type": 2}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, farms)
	})
}
