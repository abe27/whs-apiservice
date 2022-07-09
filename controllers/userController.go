package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/abe27/gin/restapi/api/v2/models"
	"github.com/abe27/gin/restapi/api/v2/services"
	databases "github.com/abe27/gin/restapi/api/v2/services"
	nanoid "github.com/aidarkhanov/nanoid/v2"
	"github.com/gin-gonic/gin"
)

var SecretKey = "admin@localhost"

func Register(c *gin.Context) {
	var r models.Response
	db := databases.DB
	user := new(models.User)
	err := c.ShouldBind(&user)
	if err != nil {
		r.Status = false
		r.Message = err.Error()
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	// Generate Nanoid ID
	id, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	user.ID = id
	user.Password = services.HashPassword(user.Password)

	/// Check Create Error
	err = db.Create(&user).Error
	if err != nil {
		r.Status = false
		r.Message = models.RegisterError
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	token, ex := services.CreateToken(user.ID)
	if ex != nil {
		r.Status = false
		r.Message = models.RegisterError
		r.Data = ex
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	// Generate Nanoid ID
	id, err = nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	var auth models.Authorization
	jwtToken := new(models.JwtToken)
	jwtToken.ID = id
	jwtToken.UserID = user.ID
	jwtToken.JwtToken = token

	err = db.Create(&jwtToken).Error
	if err != nil {
		r.Status = false
		r.Message = models.SystemError
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	auth.Type = "Bearer"
	auth.Token = jwtToken.ID
	r.Status = true
	r.Message = models.RegisterComplete
	r.Data = &auth
	c.JSON(http.StatusOK, r)
}

func Logon(c *gin.Context) {
	db := databases.DB
	var r models.Response
	var auth models.Authorization
	var login models.User

	// Parse FormData and check Validate
	err := c.ShouldBind(&login)
	if err != nil {
		r.Status = false
		r.Message = models.UserCheckFormValid
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	hand_check_passwd := login.Password
	err = db.Where("user_name=?", login.UserName).First(&login).Error
	if err != nil {
		r.Status = false
		r.Message = models.UserNotFound
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	/// Check Password Match
	err = services.ComparePassword(login.Password, hand_check_passwd)
	if err != nil {
		r.Status = false
		r.Message = models.UserPasswordNotMatch
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	id, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	var jwtToken models.JwtToken
	db.Where("user_id=?", login.ID).First(&jwtToken)
	token, e := services.CreateToken(login.ID)
	if e != nil {
		log.Fatalln(e)
	}

	jToken := new(models.JwtToken)
	jToken.ID = id
	jToken.UserID = login.ID
	jToken.JwtToken = token

	err = db.Create(&jToken).Error
	if err != nil {
		r.Status = false
		r.Message = models.UserLoginError
		r.Data = nil
		// Delete token duplicate
		db.Delete(&jwtToken)
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	auth.Token = jToken.ID
	if err != nil {
		r.Status = false
		r.Message = models.SystemError
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	auth.Type = "Bearer"
	r.Status = true
	r.Message = models.Welcome
	r.Data = &auth
	c.JSON(http.StatusOK, r)
}

func RefreshToken(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.UserRefreshToken
	r.Data = nil
	c.JSON(http.StatusOK, r)
}

func Logout(c *gin.Context) {
	db := databases.DB
	var jwtToken models.JwtToken
	var r models.Response
	s := c.Request.Header.Get("Authorization")
	jwtToken.ID = strings.TrimPrefix(s, "Bearer ")
	err := db.Find(&jwtToken).Error
	if err != nil {
		r.Status = false
		r.Message = models.SystemError
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	db.Delete(&jwtToken)
	r.Status = true
	r.Message = models.UserLeave
	r.Data = nil
	c.JSON(http.StatusOK, r)
}

func Profile(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.UserProfileReady
	r.Data = nil
	c.JSON(http.StatusOK, r)
}
