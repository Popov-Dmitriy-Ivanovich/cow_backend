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
	NameSurnamePatronimic string // ФИО
	RoleId                uint   // ID роли с которой пользователь регистрируется
	Email                 string // почта пользователя
	Phone                 string // телефон пользователя
	Password              string // пароль пользователя

	HozNumber *uint // номер хоз-ва либо существующего, либо newHoz

	RegionId uint // ID региона в котором будет зарегистрирован пользователь
}

type hozData struct {
	HozNumber  string
	DistrictId uint

	HoldNumber string // номер холдинга: либо существующего, либо newHold

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
	NewUser models.UserRegisterRequest  // данные пользователя для регистрации
	NewHoz  *models.HozRegisterRequest  // не обрабатывается
	NewHold *models.HoldRegisterRequest // не обрабатывается
}

type userClaims struct {
	jwt.RegisteredClaims
	UserData createUserData
}

// Create
// @Summary      User register request
// @Description  Рут для создания запроса на регистрацию
// @Param        userData    body     createUserData true  "applied filters"
// @Tags         User
// @Produce      json
// @Success      200  {object}   string
// @Failure      500  {object}  string
// @Failure      422  {object}  string
// @Failure      401  {object}  string
// @Router       /user/create [post]
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
		message := []byte("From: genmilk@aurusoft.ru\r\n" +
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

// VerifyEmail
// @Summary      Get list of sexes
// @Description  Ссылка должна приходить на почту пользователя автоматически, с фронтенда этот рут не фетчить
// @Tags         User
// @Produce      html
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

		c.HTML(200, "MessageResponse.tmpl", gin.H{})
	}
}
