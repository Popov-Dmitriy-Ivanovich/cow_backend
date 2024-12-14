package load

import (
	"cow_backend/routes/auth"

	"github.com/gin-gonic/gin"
)

type Load struct{}

func (s *Load) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/load")
	apiGroup.Use(auth.AuthMiddleware(auth.Farmer))
	apiGroup.POST("/cow", s.Cow())
	apiGroup.POST("/checkMilk", s.CheckMilk())
	apiGroup.POST("/event", s.Event())
	apiGroup.POST("/genetic", s.Genetic())
	apiGroup.POST("/grade", s.Grade())
	apiGroup.POST("/lactation", s.Lactation())
	apiGroup.POST("/exterior", s.Exterior())
	apiGroup.POST("/partner", s.Partner())
	apiGroup.POST("/gtcFile", s.GtcFile())
	apiGroup.POST("/document", s.Document())
}
