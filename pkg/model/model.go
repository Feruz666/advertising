package model

type PostModel struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Price       int      `json:"price"`
}
