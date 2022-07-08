package main

import (
	"github.com/abe27/gin/restapi/api/v2/databases"
	"github.com/abe27/gin/restapi/api/v2/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	databases.ConnectDatabase()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	routes.Register(r)
	r.Run(":3000")
}
