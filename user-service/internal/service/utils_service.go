package service

import(
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"user-service/internal/models"
	"os"
	"net/http"
	"time"
)

func HashPassword(c *gin.Context, hash *string, password string) int {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 10) // default cost

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return 1
	}
	*hash = string(h)
	return 0
}


func CheckPassword(c *gin.Context, user models.User, password string) int {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return 1
	}
	return 0
}

func GenerateToken(c *gin.Context, user models.User, tokenString *string) int {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return 1
	}
	*tokenString = signedToken
	return 0
}

func SetCookie(c *gin.Context, token string){
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
}