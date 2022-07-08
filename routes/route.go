package routes

import (
	"github.com/abe27/gin/restapi/api/v2/controllers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	route := r.Group("/api/v2")
	route.GET("/ping", controllers.Hello)

	// Route User interface
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Logon)
	route.GET("/profile", controllers.Profile)
	route.GET("/refresh", controllers.RefreshToken)
	route.DELETE("/logout", controllers.Logout)
}
