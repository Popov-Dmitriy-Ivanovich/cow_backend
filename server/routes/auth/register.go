package auth

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	NameSurnamePatronimic string `json:"fullname"`
	RoleID                int    `json:"role"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	Password              string `json:"password"`
	FarmId                *uint  `json:"farm"`
	RegionId              uint   `json:"region"`
}

// Register
// @Summary      REGISTER (FOR ADMIN PAGE)
// @Description  Нужно для админки
// @Tags         REGISTER
// @Param        RegisterData    body     RegisterData  true  "applied filters"
// @Accept       json
// @Produce      json
// @Success      200  {array}   string
// @Failure      422  {object}  map[string]error
// @Failure      500  {object}  map[string]error
// @Router       /auth/register [post]
func (s *Auth) Register() func(*gin.Context) {
	return func(c *gin.Context) {
		request := RegisterData{}
		db := models.GetDb()

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		existUser := models.User{}
		if err := db.Where("email = ?", request.Email).First(&existUser).Error; err == nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Пользователь с таким email уже существует"})
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

// CheckEmail checks if user with given email already exists
//
//	@Summary      CHECK EMAIL (USED FOR ADMIN PAGE)
//	@Tags         REGISTER
//	@Param        email    query     string  true  "email of user to check"
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  map[string]bool
//	@Failure      500  {object}  map[string]error
//	@Router       /auth/checkEmail [get]
func (s *Auth) CheckEmail() func(*gin.Context) {
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
