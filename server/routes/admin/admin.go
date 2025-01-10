package admin

import (
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (s *Admin) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/admin")
	apiGroup.GET("/login", s.Login())
	apiGroup.POST("/adminLogin", s.AdminLogin())
	apiGroup.GET("/adminLogout", s.AdminLogout())

	adminGroup := apiGroup.Group("")
	adminGroup.Use(AdminMiddleware())

	{
		adminGroup.GET("", s.Index())
		adminGroup.GET("/cowTable", s.CheckCowTable())
		adminGroup.GET("/checkUsers", s.CheckUsersTable())
		adminGroup.GET("/checkHoldings", s.CheckHozTable(1))
		adminGroup.GET("/checkHozs", s.CheckHozTable(2))
		adminGroup.GET("/checkFarms", s.CheckHozTable(3))
		adminGroup.GET("/checkNews", s.checkNews())
		adminGroup.GET("/createUser", s.CreateUser())
		adminGroup.GET("/createHolding", s.CreateHolding())
		adminGroup.GET("/createHoz", s.CreateHoz())
		adminGroup.GET("/createFarm", s.CreateFarm())
		adminGroup.GET("/createNews", s.CreateNews())
		adminGroup.GET("/checkEmail", s.checkEmail())
		adminGroup.GET("/userPage/:id", s.UpdateUserPage())
		adminGroup.GET("/holdingPage/:id", s.UpdateFarmPage(1))
		adminGroup.GET("/hozPage/:id", s.UpdateFarmPage(2))
		adminGroup.GET("/farmPage/:id", s.UpdateFarmPage(3))
		adminGroup.GET("/newsPage/:id", s.UpdateNewsPage())

		adminGroup.POST("/approveCows", s.ApproveCows())
		adminGroup.POST("/newUser", s.NewUser())
		adminGroup.POST("/newHolding", s.NewHolding())
		adminGroup.POST("/newHoz", s.NewHoz())
		adminGroup.POST("/newFarm", s.NewFarm())
		adminGroup.POST("/newNews", s.NewNews())

		adminGroup.DELETE("/deleteUser/:id", s.DeleteUser())
		adminGroup.DELETE("/deleteHoz/:id", s.DeleteHoz())
		adminGroup.DELETE("/deleteNews/:id", s.DeleteNews())

		adminGroup.PUT("/updateUser/:id", s.UpdateUser())
		adminGroup.PUT("/updateFarm/:id", s.UpdateFarm())
		adminGroup.PUT("/updateNews/:id", s.UpdateNews())
		adminGroup.GET("/printUser/:number", s.PrintUser())
		adminGroup.GET("/approveUser/:number", s.ApproveUser())
		adminGroup.GET("/rejectUser/:number", s.RejectUser())
	}
}
