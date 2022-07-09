package routes

import (
	"github.com/abe27/gin/restapi/api/v2/controllers"
	"github.com/abe27/gin/restapi/api/v2/services"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/ping", controllers.Hello)
	route := r.Group("/api/v2")
	// Route User interface
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Logon)
	user := r.Group("/api/v2/member", services.AuthorizationMiddleware)
	user.GET("/me", controllers.Profile)
	user.GET("/refresh", controllers.RefreshToken)
	user.DELETE("/logout", controllers.Logout)

	// Route Whs interface
	whs := r.Group("/api/v2/whs", services.AuthorizationMiddleware)
	whs.GET("/whs", controllers.Profile)
}
