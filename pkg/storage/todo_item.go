package storage

import (
	"fmt"

	"github.com/a1exander256/todo/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoItemStorage struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewTodoItemStorage(db *sqlx.DB, log *logrus.Logger) *TodoItemStorage {
	return &TodoItemStorage{
		db:  db,
		log: log,
	}
}

func (s *TodoItemStorage) Create(item models.TodoItem) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", tableTodoItems)

	row := s.db.QueryRow(query, item.Title, item.Description)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (s *TodoItemStorage) GetById(itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti WHERE ti.id = $1`, tableTodoItems)
	if err := s.db.Get(&item, query, itemId); err != nil {
		return item, err
	}
	return item, nil
}
