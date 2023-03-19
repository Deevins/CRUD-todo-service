package service

import (
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userID int, list entity.TodoList) (int, error)
	GetAll(userID int) ([]entity.TodoList, error)
	GetByID(userID, listID int) (entity.TodoList, error)
	Delete(userID, listID int) error
	Update(userID int, id int, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID int, item entity.TodoItem) (int, error)
	GetAll(userID, listID int) ([]entity.TodoItem, error)
	GetItemByID(userID, itemID int) (entity.TodoItem, error)
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
