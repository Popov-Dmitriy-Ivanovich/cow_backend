package admin

import (
	"cow_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// approveCows обрабатывает запрос на обновление поля approved для коров
func (s *Admin) ApproveCows() func(*gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			Approved    []string `json:"approved"`
			NotApproved []string `json:"notApproved"`
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

		c.JSON(http.StatusOK, gin.H{"message": "Статус коров обновлен"})
	}
}
