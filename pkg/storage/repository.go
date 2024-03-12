package storage

import (
	"database_skillfactory/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	TasksTable       = "tasks"
	UsersTable       = "users"
	TasksLablesTable = "tasks_labels"
	LabelsTable      = "labels"
)

type TaskDb struct {
	db *sqlx.DB
}

func NewTaskDb(db *sqlx.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (s *TaskDb) CreateTask(t model.Task) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (title,content) VALUES ($1,$2) RETURNING id`, TasksTable)
	rows := s.db.QueryRow(query, t.Title, t.Content)
	if err := rows.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TaskDb) GetAllTasks() ([]model.Task, error) {
	query := fmt.Sprintf("SELECT *FROM %s", TasksTable)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *TaskDb) FindByAuthor(author_id int) ([]model.Task, error) {
	query := fmt.Sprintf("SELECT id,opened,closed,author_id,assigned_id,title,content FROM %s WHERE author_id=$1", TasksTable)
	rows, err := s.db.Query(query, author_id)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (s *TaskDb) FindByTag(label_id int) ([]model.Task, error) {
	query := fmt.Sprintf("SELECT t.id,t.opened,t.closed,t.author_id,t.assigned_id,t.title,t.content FROM %s t JOIN %s l ON l.id=t.label_id WHERE l.id=$1", TasksTable, LabelsTable)
	rows, err := s.db.Query(query, label_id)
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		fmt.Println(t)
		err := rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (s *TaskDb) UpdateTask(task_id int, update model.UpdateTask) error {
	if len(update.Title) == 0 && len(update.Content) == 0 {
		return fmt.Errorf("No one updates")
	}

	if update.Content != "" && update.Title != "" {
		query := fmt.Sprintf("UPDATE %s SET title=$1, content=$2 WHERE id=$3", TasksTable)
		_, err := s.db.Exec(query, update.Title, update.Content, task_id)
		if err != nil {
			return err
		}
	}

	if update.Content == "" && update.Title != "" {
		query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", TasksTable)
		_, err := s.db.Exec(query, update.Title, task_id)
		if err != nil {
			return err
		}
	}

	if update.Content != "" && update.Title == "" {
		query := fmt.Sprintf("UPDATE %s SET content=$1 WHERE id=$2", TasksTable)
		_, err := s.db.Exec(query, update.Content, task_id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TaskDb) DeleteTask(task_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", TasksTable)
	_, err := s.db.Exec(query, task_id)
	if err != nil {
		return err
	}
	return nil
}
