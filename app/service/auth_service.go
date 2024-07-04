package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"poc-go/app/repository"
	"time"
)

type AuthService interface {
	Login(c *gin.Context)
	Validate(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func AuthServiceInit(userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (a AuthServiceImpl) Login(c *gin.Context) {
	// Get email & pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up for requested user
	user, err := a.userRepository.FindUserByEmail(body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// TODO compare password later
	// Compare sent in password with saved users password
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Invalid email or password",
	//	})
	//	return
	//}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Respond
	//c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (a AuthServiceImpl) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	// user.(models.User).Email    -->   to access specific data

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
