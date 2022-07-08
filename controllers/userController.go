package controllers

import (
	"net/http"

	"github.com/abe27/gin/restapi/api/v2/databases"
	"github.com/abe27/gin/restapi/api/v2/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var r models.Response
	db := databases.DB
	user := new(models.User)
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// email := c.PostForm("email")

	err := db.Create(&user).Error
	if err != nil {
		r.Status = false
		r.Message = models.RegisterError
		r.Data = err
		c.JSON(http.StatusInternalServerError, r)
		c.Abort()
		return
	}

	r.Status = true
	r.Message = models.RegisterComplete
	r.Data = nil
	c.JSON(http.StatusOK, r)
}

func Logon(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.Welcome
	r.Data = nil
	c.JSON(http.StatusOK, r)
	return
}

func RefreshToken(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.UserRefreshToken
	r.Data = nil
	c.JSON(http.StatusOK, r)
	return
}

func Logout(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.UserLeave
	r.Data = nil
	c.JSON(http.StatusOK, r)
	return
}

func Profile(c *gin.Context) {
	var r models.Response
	r.Status = true
	r.Message = models.UserProfileReady
	r.Data = nil
	c.JSON(http.StatusOK, r)
	return
}
