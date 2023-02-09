package service

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/pkg/repository"
)

type TodoListService struct {
	rep repository.TodoList
}

func NewTodoListService(rep repository.TodoList) *TodoListService {
	return &TodoListService{
		rep: rep,
	}
}

func (s *TodoListService) CreateList(userID int, list entity.TodoList) (int, error) {
	return s.rep.CreateList(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]entity.TodoList, error) {
	return s.rep.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (entity.TodoList, error) {
	return s.rep.GetByID(userID, listID)
}
