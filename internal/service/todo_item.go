package service

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
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

func (s *TodoItemService) Create(userID, listID int, item entity.TodoItem) (int, error) {
	_, err := s.listRep.GetByID(userID, listID)
	if err != nil {
		return 0, err
	}

	return s.rep.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]entity.TodoItem, error) {
	return s.rep.GetAll(userID, listID)
}

func (s *TodoItemService) GetItemByID(userID, itemID int) (entity.TodoItem, error) {
	return s.rep.GetItemByID(userID, itemID)
}
