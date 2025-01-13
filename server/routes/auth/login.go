package auth

import (
	"cow_backend/models"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthData struct {
	Email    string `json:"email"`    // почта
	Password string `json:"password"` // пароль, не зашифрованный
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId   uint
	RoleId   int
	RegionId uint
	FarmId   *uint
	Distid   uint
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Login
// @Summary      LOGIN
// @Description  После успешного логина возвращает словарь с ключем "token" - access token. Обеспечивает логин по JWT
// @Tags         LOGIN
// @Param        AuthData    body     AuthData  true  "applied filters"
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]error
// @Failure      401  {object}  map[string]error
// @Failure      500  {object}  map[string]error
// @Router       /auth/login [post]
func (s *Auth) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		user := AuthData{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db := models.GetDb()
		storedUser := models.User{}
		if err := db.Where("email = ?", user.Email).Preload("Farm").First(&storedUser).Error; err != nil {
			c.JSON(401, gin.H{"error": "Пользователь не найден"})
			return
		}

		if err := CheckPassword(string(storedUser.Password), user.Password); err != nil {
			c.JSON(401, gin.H{"error": "Неверный пароль"})
			return
		}

		jwtKey := os.Getenv("JWT_KEY")

		expTimeAccess := time.Now().Add(5 * time.Hour)

		claimsAccess := &JwtClaims{
			UserId:   storedUser.ID,
			RoleId:   storedUser.RoleId,
			RegionId: storedUser.RegionId,
			FarmId:   storedUser.FarmId,
			Distid:   storedUser.Farm.DistrictId,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTimeAccess),
			},
		}

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
		accessString, err := accessToken.SignedString([]byte(jwtKey))
		if err != nil {
			c.JSON(401, gin.H{"error": "ошибка создания токена"})
			return
		}
		c.JSON(200, gin.H{"token": accessString})

	}
}
