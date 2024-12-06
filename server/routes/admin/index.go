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
