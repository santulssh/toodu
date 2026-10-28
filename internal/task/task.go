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
	ErrInvalidDateStart = errors.New("Date start invalid, must be the current day or other day next day.")
	ErrInvalidDateRange = errors.New("Date range invalid, date end must be a least one day of difference from the date start.")
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Active      bool       `json:active`
	Status      string     `json:"status"`
	SubTasks    []*Task    `json:sub_tasks`
	DateStart   *time.Time `json:date_start`
	DateEnd     *time.Time `json:date_end`
	CreatedAt   time.Time  `json:created_at`
	ModifiedAt  time.Time  `json:modified_at`
}

type TaskStatus string

const (
	TaskStatusDone       TaskStatus = "DONE"
	TaskStatusUnfinished TaskStatus = "UNFINISHED"
	TaskStatusPending    TaskStatus = "PENDING"
	TaskStatusPlanning   TaskStatus = "PLANNING"
)

func NewTask(title string, description string, dateStart *time.Time, dateEnd *time.Time) (Task, error) {
	if strings.TrimSpace(title) == "" {
		return Task{}, ErrRequiredTitle
	}

	if len(title) <= 3 {
		return Task{}, ErrInvalidTitle
	}

	if dateStart != nil && dateStart.Before(time.Now()) {
		return Task{}, ErrInvalidDateStart
	}

	if (dateStart != nil && dateEnd != nil) && dateEnd.Before(*dateStart) {
		return Task{}, ErrInvalidDateRange
	}

	return Task{ID: uuid.NewV7(), Title: title, Description: description, Priority: string(TaskStatusUnfinished), Active: false, DateStart: dateStart, DateEnd: dateEnd}, nil
}

// func (t *Task) RemoveTask(id string) error {
//
// }

func (t *Task) Show() {
	print(t.Title)
}

// func (t *Task) AddSubTask(task)
// func (t *Task) RemoveSubTask(task)
