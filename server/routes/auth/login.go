package auth

import (
	"cow_backend/models"
	"crypto/sha256"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthData struct {
	Email    string
	Password string
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId uint
	Role   int
}

// ListAccounts lists all existing accounts
//
//	@Summary      LOGIN
//	@Tags         LOGIN
//	@Param        AuthData    body     AuthData  true  "applied filters"
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   string
//	@Failure      422  {object}  map[string]error
//	@Failure      500  {object}  map[string]error
//	@Router       /auth/login [post]
func (s *Auth) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		bodyData := AuthData{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, gin.H{"error": err})
				return
			}
		} else {
			c.JSON(422, gin.H{"error": "no auth data provided"})
			return
		}

		jwtKey := os.Getenv("JWT_KEY")
		expTimeAccess := time.Now().Add(5 * time.Hour)

		user := models.User{}
		db := models.GetDb()
		if err := db.First(&user, map[string]any{"email": bodyData.Email}).Error; err != nil {
			c.JSON(422, err.Error())
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(bodyData.Password))
		pasHash := hasher.Sum(nil)
		if !reflect.DeepEqual(user.Password, pasHash) {
			c.JSON(422, gin.H{"error": "wrong password"})
			return
		}

		claimsAccess := &JwtClaims{
			UserId: user.ID,
			Role:   user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTimeAccess),
			},
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
		accessString, err := accessToken.SignedString([]byte(jwtKey))
		if err != nil {
			c.JSON(500, gin.H{"error": "ошибка создания токена"})
		}
		c.JSON(200, gin.H{"token": accessString})
	}
}
