package service

import (
	"github.com/deevins/todo-restAPI/internal/entity"
	"github.com/deevins/todo-restAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
