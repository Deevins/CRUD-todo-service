package service

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
	uuid "github.com/satori/go.uuid"
)

type TodoItemService struct {
	rep     repository.TodoItem
	listRep repository.TodoList
}

func NewTodoItemService(rep repository.TodoItem, listRep repository.TodoList) *TodoItemService {
	return &TodoItemService{
		rep:     rep,
		listRep: listRep,
	}
}

func (s *TodoItemService) Create(userID, listID uuid.UUID, item entity.TodoItem) (uuid.UUID, error) {
	_, err := s.listRep.GetByID(userID, listID)
	if err != nil {
		return Nil, err
	}

	return s.rep.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID uuid.UUID) ([]entity.TodoItem, error) {
	return s.rep.GetAll(userID, listID)
}

func (s *TodoItemService) GetItemByID(userID, itemID uuid.UUID) (entity.TodoItem, error) {
	return s.rep.GetItemByID(userID, itemID)
}
