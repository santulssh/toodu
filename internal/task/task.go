package task

import (
	"errors"
	"strings"
	"time"
	"uuid"
)

var (
	ErrRequiredTitle    = errors.New("Title is required for a task.")
	ErrInvalidTitle     = errors.New("Title invalid, must be a least 3 characters.")
	ErrInvalidPriority  = errors.New("Priority invalid, must be set to LOW, MEDIUM or HIGH.")
	ErrInvalidStatus    = errors.New("Status invalid, must be set to PLANNING, PENDING, IN_PROGRESS or COMPLETED.")
	ErrInvalidDateStart = errors.New("Date start invalid, must be the current day or other day next day.")
	ErrInvalidDateRange = errors.New("Date range invalid, date end must be a least one day of difference from the date start.")
)

type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

type Status string

const (
	StatusPlanning  Status = "PLANNING"
	StatusPending   Status = "PENDING"
	StatusProgress  Status = "IN_PROGRESS"
	StatusCompleted Status = "COMPLETED"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	// SubTasks    []*Task    `json:sub_tasks`
	// DateStart   *time.Time `json:date_start`
	// DateEnd     *time.Time `json:date_end`
	CreatedAt  time.Time `json:created_at`
	ModifiedAt time.Time `json:modified_at`
}

func ValidateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return ErrRequiredTitle
	}

	if len(title) <= 3 {
		return ErrInvalidTitle
	}

	return nil
}

func ValidatePriority(priority Priority) error {
	switch priority {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return ErrInvalidPriority
	}
}

func ValidateStatus(status Status) error {
	switch status {
	case StatusCompleted, StatusPending, StatusPlanning, StatusProgress:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func NewTask(title string, description string, priority Priority, status Status) (Task, error) {
	if err := ValidateTitle(title); err != nil {
		return Task{}, err
	}

	if err := ValidatePriority(priority); err != nil {
		return Task{}, err
	}

	if err := ValidateStatus(status); err != nil {
		return Task{}, err
	}

	// 	if dateStart != nil && dateStart.Before(time.Now()) {
	// 		return Task{}, ErrInvalidDateStart
	// 	}
	//
	// 	if (dateStart != nil && dateEnd != nil) && dateEnd.Before(*dateStart) {
	// 		return Task{}, ErrInvalidDateRange
	// 	}

	return Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}, nil
}
