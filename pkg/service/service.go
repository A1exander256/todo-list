package service

import (
	"github.com/a1exander256/todo/models"
	"github.com/a1exander256/todo/pkg/storage"
	"github.com/sirupsen/logrus"
)

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	GetById(itemId int) (models.TodoItem, error)
}

type Service struct {
	TodoItem
}

func NewService(storage *storage.Storage, log *logrus.Logger) *Service {
	return &Service{
		TodoItem: NewTodoItemService(storage.TodoItem, log),
	}
}
