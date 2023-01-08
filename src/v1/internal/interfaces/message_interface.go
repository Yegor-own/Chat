package interfaces

type CreateMessage struct {
	Text string `json:"text"`
	Id   uint   `json:"id"`
}

type GetMessage struct {
	Id uint `json:"id"`
}
