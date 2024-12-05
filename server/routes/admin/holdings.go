package admin

import (
	"cow_backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CheckHoldingsTable() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		holds := []models.Farm{}
		db.Where("parrent_id is Null").Find(&holds)
		fmt.Printf("holds: %v\n", holds)
		c.HTML(http.StatusOK, "AdminHoldingsPage.tmpl", gin.H{"title": "Таблица холдингов", "holds": holds})
	}
}
