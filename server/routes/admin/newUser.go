package admin

import (
	"cow_backend/models"
	"fmt"
	"net/http"

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

		fmt.Println(request)
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}

		// user := models.User{
		// 	NameSurnamePatronimic: request.NameSurnamePatronimic,
		// 	Role:                  request.RoleID,
		// 	Email:                 request.Email,
		// 	Phone:                 request.Phone,
		// 	Password:              hashedPassword,
		// 	FarmId:                &request.FarmId,
		// }
		request.Password = string(hashedPassword)
		db := models.GetDb()

		// Сохраняем нового пользователя в базе данных
		if err := db.Create(&request).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новый пользователь добавлен"})
	}
}
