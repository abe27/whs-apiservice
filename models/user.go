package models

import "time"

type User struct {
	ID        string    `gorm:"primarykey;size:21" json:"id"`
	UserName  string    `gorm:"unique;not null;;size:10" json:"username" binding:"required"`
	Email     string    `gorm:"default:null;size:25" json:"email"`
	Password  string    `gorm:"not null;size:255" json:"password" binding:"required"`
	IsVerify  bool      `json:"is_verify" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}

type Authorization struct {
	Token string `json:"token"`
	Type  string `json:"type" default:"Bearer"`
}
