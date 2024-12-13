package admin

import (
	"cow_backend/routes/auth"

	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (s *Admin) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/admin")
	apiGroup.GET("/login", s.Login())

	adminGroup := apiGroup.Group("")
	adminGroup.Use(auth.AuthMiddleware(auth.Admin))
	adminGroup.GET("", s.Index())
	adminGroup.GET("/cowTable", s.CheckCowTable())
	adminGroup.GET("/checkUsers", s.CheckUsersTable())
	adminGroup.GET("/createUser", s.CreateUser())
	adminGroup.GET("/checkHoldings", s.CheckHozTable(1))
	adminGroup.GET("/createHolding", s.CreateHolding())
	adminGroup.GET("/checkHozs", s.CheckHozTable(2))
	adminGroup.GET("/createHoz", s.CreateHoz())
	adminGroup.GET("/checkFarms", s.CheckHozTable(3))
	adminGroup.GET("/createFarm", s.CreateFarm())
	adminGroup.GET("/checkEmail", s.checkEmail())
	adminGroup.POST("/approveCows", s.ApproveCows())
	adminGroup.POST("/newUser", s.NewUser())
	adminGroup.POST("/newHolding", s.NewHolding())
	adminGroup.POST("/newHoz", s.NewHoz())
	adminGroup.POST("/newFarm", s.NewFarm())
	adminGroup.DELETE("/deleteUser/:id", s.DeleteUser())
	adminGroup.DELETE("/deleteHoz/:id", s.DeleteHoz())
	adminGroup.GET("/userPage/:id", s.UpdateUserPage())
	adminGroup.GET("/holdingPage/:id", s.UpdateFarmPage(1))
	adminGroup.GET("/hozPage/:id", s.UpdateFarmPage(2))
	adminGroup.GET("/farmPage/:id", s.UpdateFarmPage(3))
	adminGroup.PUT("/updateUser/:id", s.UpdateUser())
	adminGroup.PUT("/updateFarm/:id", s.UpdateFarm())

}
