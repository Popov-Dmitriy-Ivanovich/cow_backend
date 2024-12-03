package gui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Gui) CowLoad() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "CowLoadPage.tmpl", gin.H{"title": "Загрузка коровы"})
	}
}
