package entities

import "time"

type Message struct {
	ID        uint      `gorm:"primary_key,autoIncrement" json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	ChatID    uint
}
