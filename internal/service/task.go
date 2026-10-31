package service

import (
	"toodu/toodu/internal/model"
	"toodu/toodu/internal/storage"
	"uuid"
)

type TaskService struct {
	repository storage.TaskRepository
}

func NewTaskService(repository storage.TaskRepository) *TaskService {
	return &TaskService{
		repository: repository,
	}
}

func (s *TaskService) CreateTask(
	title string,
	description string,
	priority string,
	status string,
) (model.Task, error) {

	newTask, err := model.NewTask(
		title,
		description,
		model.Priority(priority),
		model.Status(status),
	)
	if err != nil {
		return model.Task{}, err
	}

	if err := s.repository.Add(newTask); err != nil {
		return model.Task{}, err
	}

	return newTask, nil
}

func (s *TaskService) GetTask(id uuid.UUID) (model.Task, error) {
	task, err := s.repository.GetById(id)

	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *TaskService) ListTasks() ([]model.Task, error) {
	tasks, err := s.repository.List()
	if err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
}
