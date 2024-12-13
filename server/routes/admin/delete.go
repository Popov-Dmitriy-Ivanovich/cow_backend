package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) DeleteUser() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		user := models.User{}
		db := models.GetDb()
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь с таким ID не найден"})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})

	}

}

func (s *Admin) DeleteHoz() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		farm := models.Farm{}
		db := models.GetDb()
		if err := db.First(&farm, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Хозяйство с таким ID не найдено"})
			return
		}

		if err := db.Delete(&farm).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Хозяйство успешно удалено"})

	}

}

func (s *Admin) DeleteNews() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		news := models.News{}
		db := models.GetDb()
		if err := db.First(&news, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Новость с таким ID не найдена"})
			return
		}

		if err := db.Delete(&news).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Новость успешно удалена"})

	}
}
