package services

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/abe27/gin/restapi/api/v2/models"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/utils"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(token string) error {
	if token != "ACCESS_TOKEN" {
		return fmt.Errorf("token provided was invalid")
	}

	return nil
}

func CreateToken(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = utils.UUID()
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString([]byte(models.SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	var r models.Response
	token := strings.TrimPrefix(s, "Bearer ")

	if err := ValidateToken(token); err != nil {
		r.Status = false
		r.Message = models.UserNotAuthenticated
		r.Data = err
		c.JSON(http.StatusUnauthorized, r)
		c.Abort()
		return
	}

	c.Next()
}
