package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckUsersTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		users := []models.User{}
		db.Find(&users)
		c.HTML(http.StatusOK, "AdminUsersPage.tmpl", gin.H{"title": "Таблица пользователей", "users": users})
	}
}
