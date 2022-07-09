package models

import "time"

type JwtToken struct {
	ID        string    `json:"id"`
	UserID    string    `gorm:"unique" json:"user_id"`
	JwtToken  string    `json:"jwt_token"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}
