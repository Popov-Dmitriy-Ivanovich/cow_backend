package cows

import (
	"cow_backend/models"

	"github.com/gin-gonic/gin"
)

// Health
// @Summary      Get list of health events
// @Description  Возращает список всех ветеренарных мероприятий для конкретной коровы.
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищутся вет мероприятия"
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  string
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
		db.
			Preload("EventType").
			Preload("EventType1").
			Preload("EventType2").
			Order("date desc").
			Find(&events, "cow_id = ? AND EXISTS (SELECT 1 FROM event_types where event_types.id = events.event_type_id AND event_types.type IN (1, 2, 3, 4))", id)

		c.JSON(200, events)
	}
}
