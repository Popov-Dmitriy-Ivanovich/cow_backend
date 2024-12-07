package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateUser() func(*gin.Context) {
	return func(c *gin.Context) {

		db := models.GetDb()
		farms := []models.Farm{}
		districts := []models.District{}
		roles := []models.Role{}
		db.Where("type = 2").Find(&farms)
		db.Find(&districts)
		db.Find(&roles)

		c.HTML(http.StatusOK, "AdminCreateUserPage.tmpl", gin.H{
			"title":     "Создание пользователя",
			"farms":     farms,
			"districts": districts,
			"roles":     roles})
	}
}
