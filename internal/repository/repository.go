package repository

import (
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// Nil create zero-value for uuid to return this value in error
var Nil = uuid.UUID{}

type Authorization interface {
	CreateUser(user entity.User) (uuid.UUID, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoList interface {
	Create(userID uuid.UUID, list entity.TodoList) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]entity.TodoList, error)
	GetByID(userID, listID uuid.UUID) (entity.TodoList, error)
	Delete(userID, listID uuid.UUID) error
	Update(userID, id uuid.UUID, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(listID uuid.UUID, item entity.TodoItem) (uuid.UUID, error)
	GetAll(userID, listID uuid.UUID) ([]entity.TodoItem, error)
	GetItemByID(userID, itemID uuid.UUID) (entity.TodoItem, error)
}

// TODO: change fmt.Sprintf to sql placeholders(vs sql injections)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
