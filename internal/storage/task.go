package storage

import (
	"errors"
	"sync"
	"toodu/toodu/internal/model"
	"uuid"
)

var (
	ErrTaskNotFound      = errors.New("The search has not found the task.")
	ErrTaskIDAlreadyUsed = errors.New("This task ID is already in use.")
	ErrTasksEmpty        = errors.New("The task repository is empty.")
)

type TaskRepository interface {
	Add(task model.Task) error
	GetById(id uuid.UUID) (model.Task, error)
	List() ([]model.Task, error)
	Update(model.Task) error
	Delete(id uuid.UUID) error
}

type InMemoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[uuid.UUID]model.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[uuid.UUID]model.Task),
	}
}

func (r *InMemoryTaskRepository) Add(newTask model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exist := r.tasks[newTask.ID]; !exist {
		r.tasks[newTask.ID] = newTask
		return nil
	}

	return ErrTaskIDAlreadyUsed
}

func (r *InMemoryTaskRepository) GetById(id uuid.UUID) (model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if t, exist := r.tasks[id]; exist {
		return t, nil
	}

	return model.Task{}, ErrTaskNotFound
}

func (r *InMemoryTaskRepository) List() ([]model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.tasks) == 0 {
		return []model.Task{}, ErrTasksEmpty
	}

	tasks := make([]model.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) Update(updatedTask model.Task) error {
	r.mu.RLock()
	defer r.mu.Unlock()

	if _, exist := r.tasks[updatedTask.ID]; exist {
		r.tasks[updatedTask.ID] = updatedTask
		return nil
	}

	return ErrTaskNotFound
}

func (r *InMemoryTaskRepository) Delete(id uuid.UUID) error {
	r.mu.RLock()
	defer r.mu.Unlock()

	if _, exist := r.tasks[id]; exist {
		delete(r.tasks, id)
		return nil
	}

	return ErrTaskNotFound
}
