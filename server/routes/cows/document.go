package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Document
// @Summary      Get list of documents
// @Description  Возвращает список документов коровы
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищутся документы"
// @Produce      json
// @Success      200  {object}   []models.Document
// @Failure      500  {object}   string
// @Router       /cows/{id}/documents [get]
func (f *Cows) Document() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.
			Preload("Documents").
			First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, cow.Documents)
	}
}
