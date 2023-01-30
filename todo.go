package todo_restAPI

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id          int    `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UserLIst struct {
	Id     int
	UserID int
	ListID int
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
