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

	if err := model.ValidateTaskTitle(title); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidateTaskDescription(description); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidateTaskPriority(model.TaskPriority(priority)); err != nil {
		return model.Task{}, err
	}

	if err := model.ValidateTaskStatus(model.TaskStatus(status)); err != nil {
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

func (s *TaskService) UpdateTask(updatedTask model.Task) error {

	if err := model.ValidateTaskTitle(updatedTask.Title); err != nil {
		return err
	}

	if err := model.ValidateTaskDescription(updatedTask.Description); err != nil {
		return err
	}

	if err := model.ValidateTaskPriority(updatedTask.Priority); err != nil {
		return err
	}

	if err := model.ValidateTaskStatus(updatedTask.Status); err != nil {
		return err
	}
	err := s.repository.Update(updatedTask)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DeleteTask(id uuid.UUID) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
