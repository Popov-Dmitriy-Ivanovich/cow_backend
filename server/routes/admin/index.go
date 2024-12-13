package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminIndex.tmpl", gin.H{"title": "Меню админа"})
	}
}

func (s *Admin) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminAuthPage.tmpl", gin.H{"title": "Авторизация админа"})
	}
}
