package entities

import "time"

type User struct {
	ID        uint       `gorm:"primary_key,autoIncrement" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"hutils"`
	LastSeen  *time.Time `json:"last_seen,omitempty"`
	AvatarUrl *string    `json:"avatar_url,omitempty"`
}
