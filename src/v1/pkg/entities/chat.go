package entities

type Chat struct {
	ID    uint   `gorm:"primary_key,autoIncrement" json:"id"`
	Title string `json:"title"`
}
