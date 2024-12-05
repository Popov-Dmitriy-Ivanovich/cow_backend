package admin

import (
	"net/http"

	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (s *Admin) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/admin")
	apiGroup.GET("/cowTable", s.CheckCowTable())
	apiGroup.GET("/checkUsers", s.CheckUsersTable())
	apiGroup.GET("/checkHoldings", s.CheckHoldingsTable())
	apiGroup.GET("", s.Index())
}

func (s *Admin) CheckCowTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		cows := []models.Cow{}
		db.Where("approved = 0").Find(&cows)
		c.HTML(http.StatusOK, "AdminCowTablePage.tmpl", gin.H{"title": "Таблица коров", "cows": cows})
	}
}

func (s *Admin) CheckUsersTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		users := []models.User{}
		db.Find(&users)
		c.HTML(http.StatusOK, "AdminUsersPage.tmpl", gin.H{"title": "Таблица пользователей", "users": users})
	}
}

func (s *Admin) CheckHoldingsTable() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminHoldingsPage.tmpl", gin.H{"title": "Таблица холдингов"})
	}
}
