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
	apiGroup.GET("/checkHoldings", s.CheckHozTable(1))
	apiGroup.GET("/createHolding", s.CreateHolding())
	apiGroup.GET("/checkHozs", s.CheckHozTable(2))
	apiGroup.GET("/createHoz", s.CreateHoz())
	apiGroup.GET("/checkFarms", s.CheckHozTable(3))
	apiGroup.GET("/createFarm", s.CreateFarm())
	apiGroup.POST("/approveCows", s.ApproveCows())
	apiGroup.POST("/newUser", s.NewUser())
	apiGroup.POST("/newHolding", s.NewHolding())
	apiGroup.POST("/newHoz", s.NewHoz())
	apiGroup.POST("/newFarm", s.NewFarm())
	apiGroup.DELETE("/deleteUser/:id", s.DeleteUser())
	apiGroup.DELETE("/deleteHoz/:id", s.DeleteHoz())
	apiGroup.GET("", s.Index())
}
