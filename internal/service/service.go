package service

import (
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
	uuid "github.com/satori/go.uuid"
)

// Nil create zero-value for uuid to return this value in error
var Nil = uuid.UUID{}

type Authorization interface {
	CreateUser(user entity.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uuid.UUID, error)
}

type TodoList interface {
	Create(userID uuid.UUID, list entity.TodoList) (uuid.UUID, error)
	GetAll(userID uuid.UUID) ([]entity.TodoList, error)
	GetByID(userID, listID uuid.UUID) (entity.TodoList, error)
	Delete(userID, listID uuid.UUID) error
	Update(userID, id uuid.UUID, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID uuid.UUID, item entity.TodoItem) (uuid.UUID, error)
	GetAll(userID, listID uuid.UUID) ([]entity.TodoItem, error)
	GetItemByID(userID, itemID uuid.UUID) (entity.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
