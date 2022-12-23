package model

type Chat struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Title    string  `json:"title"`
	OwnerIds [2]uint `json:"owner_ids"`
}

func (Chat) TableName() string { return "chats" }
