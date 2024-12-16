package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Admin) UpdateUserPage() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		farms := []models.Farm{}
		regions := []models.Region{}
		roles := []models.Role{}
		db.Order("name").Where("type = 2").Find(&farms)
		db.Order("name").Find(&regions)
		db.Where("id != 4").Find(&roles)

		user := models.User{}
		if err := db.
			Preload("Role").
			Preload("Region").
			Preload("Farm").
			First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь с таким ID не найден"})
			return
		}

		c.HTML(http.StatusOK, "AdminUpdateUserPage.tmpl", gin.H{
			"title":   "Редактирование",
			"user":    user,
			"farms":   farms,
			"regions": regions,
			"roles":   roles})
	}
}

func (s *Admin) UpdateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		user := models.User{}
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь с таким ID не найден"})
			return
		}

		var request struct {
			Fullname string `json:"fullname"`
			Password string `json:"password"`
			RoleID   string `json:"role"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			RegionId string `json:"region"`
			FarmId   string `json:"farm"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if request.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
				return
			}
			user.Password = hashedPassword
		}

		if request.RoleID != "" {
			role, err := strconv.Atoi(request.RoleID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID роли"})
			}
			user.RoleId = role
		}

		if request.RegionId != "" {
			region, err := strconv.ParseUint(request.RegionId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID региона"})
			}
			regionID := uint(region)
			user.RegionId = regionID
		}

		if request.FarmId != "" {
			farm, err := strconv.ParseUint(request.FarmId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID фермы"})
			}
			farmID := uint(farm)
			user.FarmId = &farmID
		}

		if request.Fullname != "" {
			user.NameSurnamePatronimic = request.Fullname
		}

		if request.Email != "" {
			user.Email = request.Email
		}

		if request.Phone != "" {
			user.Phone = request.Phone
		}

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

}
