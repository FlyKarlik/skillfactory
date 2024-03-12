package storage

import (
	"database_skillfactory/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateTask(t model.Task) (int, error)
	GetAllTasks() ([]model.Task, error)
	FindByAuthor(author_id int) ([]model.Task, error)
	FindByTag(label_id int) ([]model.Task, error)
	UpdateTask(task_id int, update model.UpdateTask) error
	DeleteTask(task_id int) error
}

type Store struct {
	Repository
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Repository: NewTaskDb(db),
	}
}
