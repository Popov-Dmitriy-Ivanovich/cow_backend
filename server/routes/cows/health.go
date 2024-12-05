package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of health events
//	@Description  Возращает список всех ветеренарных мероприятий для конкретной коровы.
//	@Tags         Cows
//	@Param        id   path      int  true  "ID коровы для которой ищутся вет мероприятия"
//
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  map[string]error
// @Router       /cows/{id}/health [get]
func (f *Cows) Health() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDb()
		if err := db.First(&cow, id).Error; err != nil {
			c.JSON(404, err.Error())
			return
		}

		events := []models.Event{}
		db.Find(&events, "cow_id = ? AND event_type_id IN (1, 2, 3, 4)", id)

		c.JSON(200, events)
	}
}