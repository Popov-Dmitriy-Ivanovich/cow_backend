package admin

import (
	"cow_backend/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Admin) ApproveCows() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			Approved    []string  `json:"approved"`
			NotApproved []string  `json:"notApproved"`
			СurrentDate time.Time `json:"currentDate"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := models.GetDb()

		for _, idStr := range request.Approved {
			id, _ := strconv.Atoi(idStr)
			db.Model(&models.Cow{}).Where("id = ?", id).Update("approved", 1)
		}

		for _, idStr := range request.NotApproved {
			id, _ := strconv.Atoi(idStr)
			db.Model(&models.Cow{}).Where("id = ?", id).Update("approved", -1)
		}

		date := models.Update{}
		if err := db.First(&date).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		date.Date = request.СurrentDate
		log.Println(date.Date)
		if err := db.Save(&date).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Статус коров обновлен"})
	}
}
