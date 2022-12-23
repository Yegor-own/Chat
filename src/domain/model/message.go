package model

import "time"

type Message struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	AuthorID  uint      //TODO describe gorm reference
	ChatID    uint      //TODO describe gorm ref
}

func (Message) TableName() string { return "messages" }
