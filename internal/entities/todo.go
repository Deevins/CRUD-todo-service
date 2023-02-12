package entity

import "errors"

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

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
} //2:28

func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update struct has no values")
	}
	return nil
}
