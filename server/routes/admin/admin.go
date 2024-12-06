package admin

import (
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (s *Admin) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/admin")
	apiGroup.GET("/cowTable", s.CheckCowTable())
	apiGroup.GET("/checkUsers", s.CheckUsersTable())
	apiGroup.GET("/createUser", s.CreateUser())
	apiGroup.GET("/checkHoldings", s.CheckHoldingsTable())
	apiGroup.GET("/createHolding", s.CreateHolding())
	apiGroup.POST("/approveCows", s.ApproveCows())
	apiGroup.GET("", s.Index())
}
