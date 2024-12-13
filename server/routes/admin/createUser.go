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
		regions := []models.Region{}
		roles := []models.Role{}
		db.Order("name").Where("type = 2").Find(&farms)
		db.Order("name").Find(&regions)
		db.Where("id != 4").Find(&roles)

		c.HTML(http.StatusOK, "AdminCreateUserPage.tmpl", gin.H{
			"title":   "Создание пользователя",
			"farms":   farms,
			"regions": regions,
			"roles":   roles})
	}
}
