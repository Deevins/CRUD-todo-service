package service

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
	uuid "github.com/satori/go.uuid"
)

type TodoListService struct {
	rep repository.TodoList
}

func NewTodoListService(rep repository.TodoList) *TodoListService {
	return &TodoListService{
		rep: rep,
	}
}

func (s *TodoListService) Create(userID uuid.UUID, list entity.TodoList) (uuid.UUID, error) {
	return s.rep.Create(userID, list)
}

func (s *TodoListService) GetAll(userID uuid.UUID) ([]entity.TodoList, error) {
	return s.rep.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID uuid.UUID) (entity.TodoList, error) {
	return s.rep.GetByID(userID, listID)
}

func (s *TodoListService) Delete(userID, listID uuid.UUID) error {
	return s.rep.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID uuid.UUID, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.rep.Update(userID, listID, input)
}
