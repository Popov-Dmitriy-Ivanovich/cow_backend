package admin

import (
	"cow_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// /admin/approveUser/:number
func (a *Admin) ApproveUser() func(*gin.Context) {
	return func(c *gin.Context) {
		userCreateNumber := c.Param("number")
		userCreateNumberInt, err := strconv.ParseUint(userCreateNumber, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db := models.GetDb()
		userRegReq := models.UserRegisterRequest{}
		if err := db.Offset(int(userCreateNumberInt)).Limit(1).Find(&userRegReq).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegReq.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}
		hoz := models.Farm{}
		role := models.Role{}
		region := models.Region{}

		if err := db.First(&hoz, map[string]any{"hoz_number": userRegReq.HozNumber}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if err := db.First(&role, userRegReq.RoleId).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if err := db.First(&region, userRegReq.RegionId).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		userReg := models.User{
			NameSurnamePatronimic: userRegReq.NameSurnamePatronimic,
			Email:                 userRegReq.Email,
			Phone:                 userRegReq.Phone,
			Password:              hashedPassword,
			Farm:                  &hoz,
			Region:                region,
			Role:                  role,
		}
		if err := db.Create(&userReg).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if err := db.Delete(&userRegReq).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "Пользователь подтвержден")
	}
}

// /admin/rejectUser/:number
func (a *Admin) RejectUser() func(*gin.Context) {
	return func(c *gin.Context) {
		userCreateNumber := c.Param("number")
		userCreateNumberInt, err := strconv.ParseUint(userCreateNumber, 10, 64)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db := models.GetDb()
		userRegReq := models.UserRegisterRequest{}
		if err := db.Offset(int(userCreateNumberInt)).Limit(1).Find(&userRegReq).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if err := db.Delete(&userRegReq).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "Пользователь отклонен")
	}
}

// /admin/printUser/:number
func (a *Admin) PrintUser() func(*gin.Context) {
	return func(c *gin.Context) {
		userCreateNumber := c.Param("number")
		userCreateNumberInt, err := strconv.ParseUint(userCreateNumber, 10, 64)
		fmt.Println("Page printuser", userCreateNumber, userCreateNumberInt)
		if err != nil {
			c.JSON(422, err.Error())
			return
		}
		db := models.GetDb()
		userRegReq := models.UserRegisterRequest{}
		if qRes := db.Offset(int(userCreateNumberInt)).Limit(1).Find(&userRegReq); qRes.Error != nil {
			c.JSON(500, qRes.Error.Error())
			return
		} else if qRes.RowsAffected == 0 {
			c.HTML(200, "AdminApproveUserPageEnd.tmpl", gin.H{})
			return
		}

		hoz := models.Farm{}
		role := models.Role{}
		region := models.Region{}

		if err := db.First(&hoz, map[string]any{"hoz_number": userRegReq.HozNumber}).Error; err != nil {
			hoz.Name = "Ферма, указанная пользователем отсутсвует в базе данных"
			hoz.HozNumber = new(string)
			*hoz.HozNumber = "Номер фермы, указанный пользователем не найден в базе данных"
		}

		if err := db.First(&role, userRegReq.RoleId).Error; err != nil {
			role.Name = "Указанная пользователем роль не существует в базе данных"
		}

		if err := db.First(&region, userRegReq.RegionId).Error; err != nil {
			region.Name = "Указанный пользователем регион не существует в базе данных"
		}
		prevNumber := userCreateNumberInt
		if prevNumber != 0 {
			prevNumber--
		}
		params := gin.H{
			"email":      userRegReq.Email,
			"name":       userRegReq.NameSurnamePatronimic,
			"role":       role.Name,
			"phone":      userRegReq.Phone,
			"hoz":        hoz.Name,
			"hozNumber":  hoz.HozNumber,
			"region":     region.Name,
			"nextPage":   "https://genmilk.ru/api/admin/printUser/" + strconv.FormatUint(userCreateNumberInt+1, 10),
			"prevPage":   "https://genmilk.ru/api/admin/printUser/" + strconv.FormatUint(prevNumber, 10),
			"approveUrl": "https://genmilk.ru/api/admin/approveUser/" + strconv.FormatUint(userCreateNumberInt, 10),
			"rejectUrl":  "https://genmilk.ru/api/admin/rejectUser/" + strconv.FormatUint(userCreateNumberInt, 10),
		}
		c.HTML(200, "AdminApproveUserPage.tmpl", params)
	}
}
