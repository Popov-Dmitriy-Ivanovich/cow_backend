package auth

import (
	"cow_backend/models"
	"crypto/sha256"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	NameSurnamePatronimic string
	Role                  int
	Email                 string
	Phone                 string
	Password              string
	FarmId                uint
}

// ListAccounts lists all existing accounts
//
//	@Summary      REGISTER
//	@Tags         REGISTER
//	@Param        RegisterData    body     RegisterData  true  "applied filters"
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   string
//	@Failure      422  {object}  map[string]error
//	@Failure      500  {object}  map[string]error
//	@Router       /auth/register [post]
func (a *Auth) Register() func(*gin.Context) {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		bodyData := RegisterData{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, gin.H{"error": err})
				return
			}
		} else {
			c.JSON(422, gin.H{"error": "no register data provided"})
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(bodyData.Password))
		pasHash := hasher.Sum(nil)
		newUser := models.User{
			NameSurnamePatronimic: bodyData.NameSurnamePatronimic,
			RoleId:                bodyData.Role,
			Email:                 bodyData.Email,
			Phone:                 bodyData.Phone,
			FarmId:                &bodyData.FarmId,
			Password:              pasHash,
		}
		db := models.GetDb()
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "OK")
	}
}
