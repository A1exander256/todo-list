package service

import (
	"github.com/a1exander256/todo/models"
	"github.com/a1exander256/todo/pkg/storage"
	"github.com/sirupsen/logrus"
)

type TodoItemService struct {
	storage storage.TodoItem
	log     *logrus.Logger
}

func NewTodoItemService(storage storage.TodoItem, log *logrus.Logger) *TodoItemService {
	return &TodoItemService{
		storage: storage,
		log:     log,
	}
}

func (s *TodoItemService) Create(item models.TodoItem) (int, error) {
	return s.storage.Create(item)
}

func (s *TodoItemService) GetById(itemId int) (models.TodoItem, error) {
	return s.storage.GetById(itemId)
}
