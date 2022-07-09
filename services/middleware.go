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

// func ValidateToken(token string) error {
// 	db := DB
// 	var jwtToken models.JwtToken
// 	err := db.Where(&token, jwtToken.ID).First(&jwtToken).Error
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(jwtToken.JwtToken)

// 	if token != "ACCESS_TOKEN" {
// 		return fmt.Errorf("token provided was invalid")
// 	}

// 	return nil
// }

func ValidateToken(token string) (interface{}, error) {
	db := DB
	var jwtToken models.JwtToken
	err := db.Where("id=?", token).First(&jwtToken).Error
	if err != nil {
		return nil, err
	}

	// fmt.Println(token)
	// fmt.Println(jwtToken.JwtToken)
	parsedToken, err := jwt.Parse(jwtToken.JwtToken, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t)
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(models.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	fmt.Println(claims["name"])
	fmt.Println("--------------------------------------")

	return claims["name"], nil
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
	_, err := ValidateToken(token)
	if err != nil {
		r.Status = false
		r.Message = err.Error()
		r.Data = err
		c.JSON(http.StatusUnauthorized, r)
		c.Abort()
		return
	}
	c.Next()
}
