package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"toodu/toodu/internal/model"
	"uuid"
)

var (
	ErrTaskNotFound      = errors.New("The search has not found the task.")
	ErrTaskIDAlreadyUsed = errors.New("This task ID is already in use.")
	ErrTasksEmpty        = errors.New("The task repository is empty.")
)

type SQLiteTaskRepository struct {
	db *sql.DB
}

func NewSQLiteTaskRepository(db *sql.DB) *SQLiteTaskRepository {
	return &SQLiteTaskRepository{db: db}
}

const (
	queryAddTask = `
	INSERT INTO tasks (id, title, description, priority, status, created_at, modified_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	queryGetTaskByID = `
	SELECT id, title, description, priority, status, created_at, modified_at FROM tasks
	WHERE id = ?
	`

	queryGetTaskList = `
	SELECT id, title, description, priority, status, created_at, modified_at FROM tasks
	`

	queryUpdateTask = `
	UPDATE tasks
	SET title = ?, description = ?, priority = ?, status = ?, modified_at = ?
	WHERE id = ?
	`

	queryDeleteTask = `
	DELETE FROM tasks WHERE id = ?
	`
)

func (r *SQLiteTaskRepository) Add(newTask model.Task) error {
	_, err := r.db.Exec(
		queryAddTask,
		newTask.ID.String(),
		newTask.Title,
		newTask.Description,
		newTask.Priority,
		newTask.Status,
		newTask.CreatedAt,
		newTask.ModifiedAt,
	)
	if err != nil {
		return fmt.Errorf(ErrTaskIDAlreadyUsed.Error()+":\n\n%w", err)
	}

	return nil
}

func (r *SQLiteTaskRepository) GetById(id uuid.UUID) (model.Task, error) {
	row := r.db.QueryRow(queryGetTaskByID, id.String())

	var t model.Task
	var idStr string
	var createdAtStr, modifedAtStr string

	err := row.Scan(&idStr, &t.Title, &t.Description, &t.Priority, &t.Status, &t.CreatedAt, &t.ModifiedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, ErrTaskNotFound
		}
		return model.Task{}, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return model.Task{}, err
	}
	t.ID = parsedID
	t.CreatedAt, _ = time.Parse(time.RFC3339, createdAtStr)
	t.ModifiedAt, _ = time.Parse(time.RFC3339, modifedAtStr)

	return t, nil
}

func (r *SQLiteTaskRepository) List() ([]model.Task, error) {
	rows, err := r.db.Query(queryGetTaskList)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		var idStr string
		var createdAtStr, modifiedAtStr string

		if err := rows.Scan(&idStr, &t.Title, &t.Description, &t.Priority, &t.Status, &createdAtStr, &modifiedAtStr); err != nil {
			return nil, err
		}

		parsedID, err := uuid.Parse(idStr)
		if err != nil {
			return []model.Task{}, err
		}
		t.ID = parsedID
		t.CreatedAt, _ = time.Parse(time.RFC3339, createdAtStr)
		t.ModifiedAt, _ = time.Parse(time.RFC3339, modifiedAtStr)

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return []model.Task{}, ErrTasksEmpty
	}

	return tasks, nil
}

func (r *SQLiteTaskRepository) Update(updatedTask model.Task) error {
	result, err := r.db.Exec(
		queryUpdateTask,
		updatedTask.Title,
		updatedTask.Description,
		updatedTask.Priority,
		updatedTask.Status,
		updatedTask.ModifiedAt,
		updatedTask.ID.String(),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *SQLiteTaskRepository) Delete(id uuid.UUID) error {
	result, err := r.db.Exec(queryDeleteTask, id.String())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}
