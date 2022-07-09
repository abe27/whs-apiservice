package main

import (
	"github.com/abe27/gin/restapi/api/v2/routes"
	"github.com/abe27/gin/restapi/api/v2/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	services.ConnDB()
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	routes.Register(r)
	r.Run(":3000")
}
