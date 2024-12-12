package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Admin) NewUser() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			NameSurnamePatronimic string `json:"fullname"`
			RoleID                int    `json:"role"`
			Email                 string `json:"email"`
			Phone                 string `json:"phone"`
			Password              string `json:"password"`
			FarmId                *uint  `json:"farm"`
			RegionId              uint   `json:"region"`
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

		user := models.User{
			NameSurnamePatronimic: request.NameSurnamePatronimic,
			RoleId:                request.RoleID,
			Email:                 request.Email,
			Phone:                 request.Phone,
			Password:              hashedPassword,
			FarmId:                request.FarmId,
			RegionId:              request.RegionId,
		}
		db := models.GetDb()

		if err := updateSequenceUser(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении последовательности: " + err.Error()})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя: " + err.Error()})
			return
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

func (s *Admin) checkEmail() func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")

		db := models.GetDb()
		user := models.User{}

		if err := db.Where("email = ?", email).First(&user).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"exists": true})
			return
		}

		c.JSON(http.StatusOK, gin.H{"exists": false})
	}

}
