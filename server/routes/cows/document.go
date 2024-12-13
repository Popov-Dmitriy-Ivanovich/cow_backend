package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of check milks
//	@Description  Возращает генетическую информацию для коровы, null, если нет
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищется генетическая информация"
//
// @Produce      json
// @Success      200  {object}   []models.Document
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/documents [get]
func (f *Cows) Document() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.
			Preload("Documents").
			First(&cow, id).Error; err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, cow.Documents)
	}
}
