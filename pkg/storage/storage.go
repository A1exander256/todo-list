package storage

import (
	"github.com/a1exander256/todo/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	GetById(itemId int) (models.TodoItem, error)
}

type Storage struct {
	TodoItem
}

func NewStorage(db *sqlx.DB, log *logrus.Logger) *Storage {
	return &Storage{
		TodoItem: NewTodoItemStorage(db, log),
	}
}
