package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateHolding() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AdminCreateHoldingPage.tmpl", gin.H{"title": "Создание холдинга"})
	}
}