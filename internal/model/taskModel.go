package model

import (
	"errors"
	"strings"
	"time"
	"uuid"
)

var (
	ErrRequiredTaskTitle      = errors.New("The title is required and must be provided.")
	ErrTaskTitleTooShort      = errors.New("The title is invalid and must be at least three characters long.")
	ErrTaskTitleTooLong       = errors.New("The title is invalid and must be be less than 50 characters.")
	ErrTaskDescriptionTooLong = errors.New("The descriptions is invalid and must be be less than 200 characters.")
	ErrInvalidTaskPriority    = errors.New("Priority invalid, must be set to LOW, MEDIUM or HIGH.")
	ErrInvalidTaskStatus      = errors.New("Status invalid, must be set to PLANNING, PENDING, IN_PROGRESS or COMPLETED.")
	ErrInvalidTaskDateStart   = errors.New("Date start invalid, must be the current day or other day next day.")
	ErrInvalidTaskDateRange   = errors.New("Date range invalid, date end must be a least one day of difference from the date start.")
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "LOW"
	TaskPriorityMedium TaskPriority = "MEDIUM"
	TaskPriorityHigh   TaskPriority = "HIGH"
)

type TaskStatus string

const (
	TaskStatusPlanning  TaskStatus = "PLANNING"
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusProgress  TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted TaskStatus = "COMPLETED"
)

type Task struct {
	ID          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	// SubTasks    []*Task    `json:sub_tasks`
	// DateStart   *time.Time `json:date_start`
	// DateEnd     *time.Time `json:date_end`
	CreatedAt  time.Time `json:created_at`
	ModifiedAt time.Time `json:modified_at`
}

func ValidateTaskTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return ErrRequiredTaskTitle
	}

	if len(title) <= 3 {
		return ErrTaskTitleTooShort
	}

	if len(title) > 50 {
		return ErrTaskTitleTooLong
	}

	return nil
}

func ValidateTaskDescription(description string) error {
	if len(description) > 200 {
		return ErrTaskDescriptionTooLong
	}

	return nil
}

func ValidateTaskPriority(priority TaskPriority) error {
	switch priority {
	case TaskPriorityLow, TaskPriorityMedium, TaskPriorityHigh:
		return nil
	default:
		return ErrInvalidTaskPriority
	}
}

func ValidateTaskStatus(status TaskStatus) error {
	switch status {
	case TaskStatusCompleted, TaskStatusPending, TaskStatusPlanning, TaskStatusProgress:
		return nil
	default:
		return ErrInvalidTaskStatus
	}
}

func NewTask(title string, description string, priority string, status string) (Task, error) {
	return Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Priority:    TaskPriority(priority),
		Status:      TaskStatus(status),
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}, nil
}
