package storage

import (
	"errors"
	"testing"
	"time"
	"uuid"

	"toodu/toodu/internal/database"
	"toodu/toodu/internal/model"
)

func setupTestDB(t *testing.T) *SQLiteTaskRepository {
	db, err := database.NewSQLiteConnection(":memory:")
	if err != nil {
		t.Fatalf("failed to create test db: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return NewSQLiteTaskRepository(db)
}

func TestSQLiteTaskRepository_AddAndGetById(t *testing.T) {
	repo := setupTestDB(t)

	newTask := model.Task{
		ID:          uuid.New(),
		Title:       "Test Task SQLite",
		Description: "Testing SQLite integration",
		Priority:    model.TaskPriorityHigh,
		Status:      model.TaskStatusPending,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	err := repo.Add(newTask)
	if err != nil {
		t.Fatalf("expected no error on Add, got %v", err)
	}

	fetchedTask, err := repo.GetById(newTask.ID)
	if err != nil {
		t.Fatalf("expected no error on GetById, got %v", err)
	}

	if fetchedTask.Title != newTask.Title {
		t.Errorf("expected title %q, got %q", newTask.Title, fetchedTask.Title)
	}
}

func TestSQLiteTaskRepository_NotFound(t *testing.T) {
	repo := setupTestDB(t)

	_, err := repo.GetById(uuid.New())
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("expected error %v, got %v", ErrTaskNotFound, err)
	}
}
