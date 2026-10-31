package model

import (
	"errors"
	"strings"
	"time"
	"uuid"
)

var (
	ErrRequiredTitle      = errors.New("The title is required and must be provided.")
	ErrTitleTooShort      = errors.New("The title is invalid and must be at least three characters long.")
	ErrTitleTooLong       = errors.New("The title is invalid and must be be less than 50 characters.")
	ErrDescriptionTooLong = errors.New("The descriptions is invalid and must be be less than 200 characters.")
	ErrInvalidPriority    = errors.New("Priority invalid, must be set to LOW, MEDIUM or HIGH.")
	ErrInvalidStatus      = errors.New("Status invalid, must be set to PLANNING, PENDING, IN_PROGRESS or COMPLETED.")
	ErrInvalidDateStart   = errors.New("Date start invalid, must be the current day or other day next day.")
	ErrInvalidDateRange   = errors.New("Date range invalid, date end must be a least one day of difference from the date start.")
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
		return ErrTitleTooShort
	}

	if len(title) > 50 {
		return ErrTitleTooLong
	}

	return nil
}

func ValidateDescription(description string) error {
	if len(description) > 200 {
		return ErrDescriptionTooLong
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

	if err := ValidateDescription(title); err != nil {
		return Task{}, err
	}

	if err := ValidatePriority(priority); err != nil {
		return Task{}, err
	}

	if err := ValidateStatus(status); err != nil {
		return Task{}, err
	}

	return Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      status,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}, nil
}
