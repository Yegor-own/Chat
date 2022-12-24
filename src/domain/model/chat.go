package model

type Chat struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	Title         string `json:"title"`
	CreatorId     uint   //TODO describe gorm ref
	ParticipantId uint   //TODO describe gorm ref
}

func (Chat) TableName() string { return "chats" }
