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
	apiGroup.GET("", s.Index())
	apiGroup.GET("/cowTable", s.CheckCowTable())
	apiGroup.GET("/checkUsers", s.CheckUsersTable())
	apiGroup.GET("/createUser", s.CreateUser())
	apiGroup.GET("/checkHoldings", s.CheckHozTable(1))
	apiGroup.GET("/createHolding", s.CreateHolding())
	apiGroup.GET("/checkHozs", s.CheckHozTable(2))
	apiGroup.GET("/createHoz", s.CreateHoz())
	apiGroup.GET("/checkFarms", s.CheckHozTable(3))
	apiGroup.GET("/createFarm", s.CreateFarm())
	apiGroup.GET("/checkEmail", s.checkEmail())
	apiGroup.GET("/checkNews", s.checkNews())
	apiGroup.GET("/createNews", s.CreateNews())

	apiGroup.POST("/approveCows", s.ApproveCows())
	apiGroup.POST("/newUser", s.NewUser())
	apiGroup.POST("/newHolding", s.NewHolding())
	apiGroup.POST("/newHoz", s.NewHoz())
	apiGroup.POST("/newFarm", s.NewFarm())
	apiGroup.POST("/newNews", s.NewNews())

	apiGroup.DELETE("/deleteUser/:id", s.DeleteUser())
	apiGroup.DELETE("/deleteHoz/:id", s.DeleteHoz())
	apiGroup.DELETE("/deleteNews/:id", s.DeleteNews())

	apiGroup.GET("/userPage/:id", s.UpdateUserPage())
	apiGroup.GET("/holdingPage/:id", s.UpdateFarmPage(1))
	apiGroup.GET("/hozPage/:id", s.UpdateFarmPage(2))
	apiGroup.GET("/farmPage/:id", s.UpdateFarmPage(3))
	apiGroup.GET("/newsPage/:id", s.UpdateNewsPage())

	apiGroup.PUT("/updateUser/:id", s.UpdateUser())
	apiGroup.PUT("/updateFarm/:id", s.UpdateFarm())
	apiGroup.PUT("/updateNews/:id", s.UpdateNews())
}
