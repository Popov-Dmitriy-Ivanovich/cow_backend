package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminCreateUserPage.tmpl", gin.H{"title": "Создание пользователя"})
	}
}
