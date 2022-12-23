package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	LastSeen  time.Time `json:"last_seen"`
	AvatarUrl string    `json:"avatar_url"`
}

func (User) TableName() string { return "users" }
