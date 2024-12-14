package admin

import (
	"cow_backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := sessions.NewCookieStore([]byte(os.Getenv("JWT_KEY")))
		session, _ := store.Get(c.Request, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.Redirect(http.StatusFound, "/api/admin/login")
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *Admin) AdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := sessions.NewCookieStore([]byte(os.Getenv("JWT_KEY")))
		session, _ := store.Get(c.Request, "cookie-name")

		var user AuthData

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := models.GetDb()
		storedUser := models.User{}

		if err := db.Where("email = ?", user.Email).First(&storedUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			return
		}

		if err := CheckPassword(string(storedUser.Password), user.Password); err != nil {
			c.JSON(401, gin.H{"error": "Неверный пароль"})
			return
		}

		if storedUser.RoleId != 4 {
			c.JSON(401, gin.H{"error": "Не администратор"})
			return
		}

		session.Values["authenticated"] = true

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   10 * 3600,
			HttpOnly: true,
		}

		if err := session.Save(c.Request, c.Writer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения сессии"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Успешный вход"})
	}
}

func (s *Admin) AdminLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := sessions.NewCookieStore([]byte(os.Getenv("JWT_KEY")))
		session, _ := store.Get(c.Request, "cookie-name")

		session.Values["authenticated"] = false

		if err := session.Save(c.Request, c.Writer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения сессии"}) // 500 Internal Server Error
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вышли из системы"}) // 200 OK
	}
}
