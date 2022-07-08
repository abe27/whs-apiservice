package controllers

import (
	"net/http"

	"github.com/abe27/gin/restapi/api/v2/models"
	"github.com/gin-gonic/gin"
)

var r models.Response

func Hello(c *gin.Context) {
	r.Status = true
	r.Message = "สวัสดี! ยินดีต้อนรับเข้าสู่ระบบ Api By Gin Framework"
	r.Data = nil
	c.JSON(http.StatusOK, r)
}
