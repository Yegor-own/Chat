package model

import "time"

type User struct {
	ID        uint       `gorm:"primary_key,autoIncrement" json:"id"`
	Name      string     `gorm:"primary_key" json:"name"`
	Password  string     `json:"password"`
	LastSeen  *time.Time `json:"last_seen,omitempty"`
	AvatarUrl *string    `json:"avatar_url,omitempty"`
}

func (User) TableName() string { return "users" }
