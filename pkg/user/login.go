package user

import (
	"divviup-client/pkg/common/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler)  Login(c *gin.Context) {

		var authInput AuthInput

		if err := c.ShouldBindJSON(&authInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userFound models.User
		h.DB.Where("username=?", authInput.Username).Find(&userFound)

		if userFound.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}

		generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  userFound.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := generateToken.SignedString([]byte(viper.Get("JWT_SECRET").(string)))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		}

		c.JSON(200, gin.H{
			"token": token,
		})
}