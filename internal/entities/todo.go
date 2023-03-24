package entity

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type TodoList struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
}

type TodoItem struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	Done        bool      `json:"done" db:"done"`
}

type UserList struct {
	Id     uuid.UUID
	UserID uuid.UUID
	ListID uuid.UUID
}

type ListItem struct {
	Id     uuid.UUID
	ListId uuid.UUID
	ItemId uuid.UUID
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update struct has no values")
	}
	return nil
}
