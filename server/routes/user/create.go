package user_create

import (
	"cow_backend/models"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type userData struct {
	NameSurnamePatronimic string
	RoleId                uint
	Email                 string
	Phone                 string
	Password              string

	HozNumber 			  *uint // номер хоз-ва либо существующего, либо newHoz

	RegionId              uint
}

type hozData struct {
	HozNumber   string
	DistrictId  uint
	
	HoldNumber 	string // номер холдинга: либо существующего, либо newHold
	
	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}

type holdData struct {
	HozNumber   string
	DistrictId  string
	Name        string
	ShortName   string
	Inn         *string
	Address     *string
	Phone       *string
	Email       *string
	Description *string
	CowsCount   *uint
}

type createUserData struct {
	NewUser models.UserRegisterRequest
	NewHoz  *models.HozRegisterRequest
	NewHold *models.HoldRegisterRequest
}

type userClaims struct {
	jwt.RegisteredClaims
	UserData createUserData
}

// ListAccounts lists all existing accounts
//
//	@Summary      Get list of sexes
//	@Description  Возращает список полов
//	@Param        filter    body     createUserData true  "applied filters"
//	@Tags         User
//	@Produce      json
//	@Success      200  {array}   createUserData
//	@Failure      500  {object}  string
//	@Router       /user/create [post]
func (u *User) Create() func(*gin.Context) {
	return func(c *gin.Context) {
		userData := createUserData{}
		if err := c.ShouldBindJSON(&userData); err != nil {
			c.JSON(422, err.Error())
			return
		}

		db := models.GetDb()
		sameCount := int64(0)
		if err := db.Model(&models.User{}).Where("email = ?", userData.NewUser.Email).Count(&sameCount).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		if sameCount != 0 {
			c.JSON(422, "User already exist")
			return
		}

		jwtKey := os.Getenv("JWT_KEY")

		expTimeAccess := time.Now().Add(1 * time.Hour)
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
			UserData: userData,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTimeAccess),
			},
		})
		accessString, err := accessToken.SignedString([]byte(jwtKey))
		if err != nil {
			c.JSON(401, gin.H{"error": "ошибка создания токена"})
			return
		}
		from := os.Getenv("EMAIL_FROM")
		password := os.Getenv("EMAIL_PASS")
		to := []string{userData.NewUser.Email}
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := os.Getenv("SMTP_PORT")
		message := []byte("From: genmilk genmilk@aurusoft.ru\r\n"+	
			"To: " + userData.NewUser.Email + "\r\n" +
				"Subject: Подтвердите эл. почту\r\n" +
				"\r\n" +
				"Для подтверждения почты перейдите по ссылке: https://genmilk.ru/api/user/verifyEmail?data=" + accessString + " .\r\n")
		auth := smtp.PlainAuth("", from, password, smtpHost)
		fmt.Println(from, password)
		if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, "Email Sent Successfully!")
	}
}

// @Summary      Get list of sexes
// @Description  Запрос на валидацию имэйла
// @Tags         User
// @Produce      json
// @Success      200  {array}   models.Sex
// @Failure      500  {object}  map[string]error
// @Router       /user/verifyEmail [get]
func (u *User) VerifyEmail() func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.Query("data")
		userClaims := &userClaims{}

		token, err := jwt.ParseWithClaims(data, userClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(422, "ошибка подтверждения:"+err.Error())
			return
		}
		
		db := models.GetDb()
		newUser := userClaims.UserData.NewUser
		newHold := userClaims.UserData.NewHold
		newHoz := userClaims.UserData.NewHoz
		

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		if newHold != nil {
			if err := db.Create(newHold).Error; err != nil {
				c.JSON(500, err.Error())
				return
			}
		}
		if newHoz != nil {
			if err := db.Create(newHold).Error; err != nil {
				c.JSON(500, err.Error())
				return
			}
		}
		
		
		c.JSON(200, "Запрос отправлен на рассмотрение")
	}
}
