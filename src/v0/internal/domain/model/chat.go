package model

type Chat struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
}

func (Chat) TableName() string { return "chats" }
