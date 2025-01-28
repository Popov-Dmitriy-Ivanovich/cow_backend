package farms

import (
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
	apiGroup.GET("/hoz", f.GetHoz())
	apiGroup.GET("/hold", f.GetHoldings())
	apiGroup.GET("/farm", f.GetFarms())
}
