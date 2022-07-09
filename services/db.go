package services

import (
	"os"
	"strings"

	"github.com/abe27/gin/restapi/api/v2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnDB() {
	dsn := os.Getenv("DSN_STRING") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate database schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.JwtToken{})
	DB = db
}
