package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Admin) NewUser() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			NameSurnamePatronimic string `json:"fullname"`
			RoleID                string `json:"role"`
			Email                 string `json:"email"`
			Phone                 string `json:"phone"`
			Password              string `json:"password"`
			FarmId                string `json:"farm"`
			RegionId              string `json:"region"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}

		role, err := strconv.Atoi(request.RoleID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID роли"})
		}
		region, err := strconv.ParseUint(request.RegionId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID региона"})
		}
		regionID := uint(region)
		farm, err := strconv.ParseUint(request.FarmId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID фермы"})
		}
		farmID := uint(farm)

		user := models.User{
			NameSurnamePatronimic: request.NameSurnamePatronimic,
			RoleId:                int(role),
			Email:                 request.Email,
			Phone:                 request.Phone,
			Password:              hashedPassword,
			FarmId:                &farmID,
			RegionId:              regionID,
		}
		db := models.GetDb()

		// Сохраняем нового пользователя в базе данных
		if err := db.Create(&user).Error; err != nil {
			if err := updateSequenceUser(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении последовательности: " + err.Error()})
				return
			}

			if err := db.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя после обновления последовательности: " + err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новый пользователь добавлен"})
	}
}

func updateSequenceUser() error {
	var maxID uint
	db := models.GetDb()
	if err := db.Model(&models.User{}).Select("max(id)").Scan(&maxID).Error; err != nil {
		return err
	}

	if err := db.Exec("SELECT setval(pg_get_serial_sequence('users', 'id'), ?)", maxID).Error; err != nil {
		return err
	}
	return nil
}
