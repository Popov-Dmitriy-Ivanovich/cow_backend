package user_create

import (
	"cow_backend/models"
	"github.com/gin-gonic/gin"
)

// Whoami
// @Summary      Получить информацию о пользователе
// @Description  Рут вернет данные о пользователе из БД
// @Tags         User
// @Produce      json
// @Success      200  {object}   string
// @Failure      500  {object}  string
// @Failure      401  {object}  string
// @Router       /user/whoami [get]
func (u *User) Whoami() func(*gin.Context) {
	return func(c *gin.Context) {
		userId, ok := c.Get("UserId")
		if !ok {
			c.JSON(401, "Юля, пожалуйста почини токен!")
			return
		}
		db := models.GetDb()
		user := models.User{}
		if err := db.First(&user, userId).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, user)
	}
}
