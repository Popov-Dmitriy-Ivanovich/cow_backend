package gui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Gui) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "Index.tmpl", gin.H{"title": "Меню"})
	}
}
