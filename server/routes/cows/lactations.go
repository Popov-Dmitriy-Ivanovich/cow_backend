package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of lactations
//	@Description  Возращает список всех лактаций для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся лактации"
//
// @Produce      json
// @Success      200  {array}   models.Lactation
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/lactations [get]
func (f *Cows) Lactations() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Lactation").First(&cow, id).Error; err != nil {
			c.JSON(500, err)
			return
		}

		c.JSON(200, cow.Lactation)
	}
}
