package admin

import (
	"cow_backend/models"
	"fmt"
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
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(request)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}

		role, err := strconv.Atoi(request.RoleID)
		farm, err := strconv.ParseUint(request.FarmId, 10, 64)
		farmID := uint(farm)

		user := models.User{
			NameSurnamePatronimic: request.NameSurnamePatronimic,
			Role:                  int(role),
			Email:                 request.Email,
			Phone:                 request.Phone,
			Password:              hashedPassword,
			FarmId:                &farmID,
		}
		db := models.GetDb()

		// Сохраняем нового пользователя в базе данных
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новый пользователь добавлен"})
	}
}
