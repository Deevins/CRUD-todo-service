package entity

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UserList struct {
	Id     int
	UserID int
	ListID int
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
