package repository

import (
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoList interface {
	CreateList(userID int, list entity.TodoList) (int, error)
	GetAll(userID int) ([]entity.TodoList, error)
	GetByID(userID, listID int) (entity.TodoList, error)
	Delete(userID, listID int) error
	Update(userID int, id int, input entity.UpdateListInput) error
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      nil,
	}
}
