package service

import (
	"toodu/toodu/internal/model"
	"uuid"
)

type TaskRepository interface {
	Add(task model.Task) error
	GetById(id uuid.UUID) (model.Task, error)
	List() ([]model.Task, error)
	Update(model.Task) error
	Delete(id uuid.UUID) error
}

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
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

	if err := model.ValidateTitle(title); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidateDescription(description); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidatePriority(model.Priority(priority)); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidateStatus(model.Status(status)); err != nil {
		return model.Task{}, err
	}

	newTask, err := model.NewTask(
		title,
		description,
		priority,
		status,
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
