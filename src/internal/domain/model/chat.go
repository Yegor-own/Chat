package model

type Chat struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Users [2]User
	//CreatorId     uint   //TODO describe gorm ref
	//ParticipantId uint   //TODO describe gorm ref
}

func (Chat) TableName() string { return "chats" }
