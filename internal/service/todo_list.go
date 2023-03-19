package service

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
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

func (s *TodoListService) Delete(userID, listID int) error {
	return s.rep.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.rep.Update(userID, listID, input)
}
