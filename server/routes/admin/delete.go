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

func (s *Admin) DeleteCascade() func(*gin.Context) {
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

		if farm.Type == 1 {
			var Hozs []models.Farm
			if err := db.Where("parrent_id = ? AND type = ?", farm.ID, 2).Find(&Hozs).Error; err == nil {
				for _, child := range Hozs {
					// Удаляем дочерние фермы типа 3
					var Farms []models.Farm
					db.Where("parrent_id = ? AND type = ?", child.ID, 3).Find(&Farms)
					for _, childFarm := range Farms {
						if err := db.Delete(&childFarm).Error; err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении дочерней фермы: " + err.Error()})
							return
						}
					}

					// Удаляем само дочернее хозяйство
					if err := db.Delete(&child).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении дочернего хозяйства: " + err.Error()})
						return
					}
				}
			}

		}

		// После удаления всех дочерних Хозяйств удаляем само Холдинг
		if err := db.Delete(&farm).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Хозяйство успешно удалено"})

	}

}
