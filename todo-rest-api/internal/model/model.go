package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}
