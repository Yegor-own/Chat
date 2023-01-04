package interfaces

type IdChat struct {
	Id uint `json:"id"`
}

type CreateChat struct {
	Title string `json:"title"`
}

type UpdateChat struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}
