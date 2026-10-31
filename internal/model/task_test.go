package model

import (
	"errors"
	"testing"
	"time"
	"uuid"
)

func TestCreateTask(t *testing.T) {
	title := "Title example"
	description := "Description example"
	priority := "LOW"
	status := "PENDING"

	t.Run("Should return error when creating task with invalid title", func(t *testing.T) {
		_, err := NewTask("", description, priority, status)
		if !errors.Is(err, ErrRequiredTitle) {
			t.Errorf("Expected error %v, got %v", ErrRequiredTitle, err)
		}
	})

	t.Run("Should return error when creating task with invalid description", func(t *testing.T) {
		description := string(make([]byte, 201))
		_, err := NewTask(title, description, priority, status)
		if !errors.Is(err, ErrDescriptionTooLong) {
			t.Errorf("Expected error %v, got %v", ErrDescriptionTooLong, err)
		}
	})

	t.Run("Should return error when creating task with invalid priority", func(t *testing.T) {
		_, err := NewTask(title, description, "INVALID", status)
		if !errors.Is(err, ErrInvalidPriority) {
			t.Errorf("Expected error %v, got %v", ErrInvalidPriority, err)
		}
	})

	t.Run("Should return error when creating task with invalid status", func(t *testing.T) {
		_, err := NewTask(title, description, priority, "INVALID")
		if !errors.Is(err, ErrInvalidStatus) {
			t.Errorf("Expected error %v, got %v", ErrInvalidStatus, err)
		}
	})

	t.Run("Should create task with default status and filled date", func(t *testing.T) {
		beforeCreation := time.Now()
		task, err := NewTask(title, description, priority, status)
		afterCreation := time.Now()

		if err != nil {
			t.Errorf("Expected: created task, got %v", err)
		}

		if task.Title != title {
			t.Errorf("Expected title %s, got %s", title, task.Title)
		}

		if task.Description != description {
			t.Errorf("Expected description %s, got %s", description, task.Description)
		}

		if string(task.Priority) != priority {
			t.Errorf("Expected title %s, got %s", priority, task.Priority)
		}

		if string(task.Status) != status {
			t.Errorf("Expected title %s, got %s", status, task.Status)
		}

		// 3. Valida se o ID foi gerado e não é nulo/vazio
		if task.ID == uuid.Nil() {
			t.Error("Expected a valid UUID, got uuid.Nil")
		}

		// 4. Valida se CreatedAt e ModifiedAt estão dentro do intervalo de tempo da execução
		if task.CreatedAt.Before(beforeCreation) || task.CreatedAt.After(afterCreation) {
			t.Errorf("CreatedAt %v is out of expected time range", task.CreatedAt)
		}
		if task.ModifiedAt.Before(beforeCreation) || task.ModifiedAt.After(afterCreation) {
			t.Errorf("ModifiedAt %v is out of expected time range", task.ModifiedAt)
		}
	})
}

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		expectedErr error
	}{
		{
			name:        "Should accept a title ",
			title:       "Title example",
			expectedErr: nil,
		},
		{
			name:        "Should not permit empty title",
			title:       "",
			expectedErr: ErrRequiredTitle,
		},
		{
			name:        "Should not permit title with 3 characters or less",
			title:       string(make([]byte, 3)),
			expectedErr: ErrTitleTooShort,
		},
		{
			name:        "Should not permit title with more then 50 characters",
			title:       string(make([]byte, 51)),
			expectedErr: ErrTitleTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTitle(tt.title)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ValidateTitle() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		expectedErr error
	}{
		{
			name:        "Should accept a description",
			description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus bibendum sagittis est pretium fringilla. Nam ultricies blandit enim, sit amet aliquet purus vestibulum sed. In finibus eleifend sapien.",
			expectedErr: nil,
		},
		{
			name:        "Should accept empty description",
			description: "",
			expectedErr: nil,
		},
		{
			name:        "Should not permit descriptions with more then 200 characters",
			description: string(make([]byte, 201)),
			expectedErr: ErrDescriptionTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDescription(tt.description)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ValidateDescription() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidatePriority(t *testing.T) {
	tests := []struct {
		name        string
		priority    string
		expectedErr error
	}{
		{
			name:        "Should accept LOW priority",
			priority:    "LOW",
			expectedErr: nil,
		},
		{
			name:        "Should accept MEDIUM priority",
			priority:    "MEDIUM",
			expectedErr: nil,
		},
		{
			name:        "Should accept HIGH priority",
			priority:    "HIGH",
			expectedErr: nil,
		},
		{
			name:        "Should return Invalid Priority",
			priority:    "URGENT",
			expectedErr: ErrInvalidPriority,
		}, {
			name:        "Should return Error Empty Priority",
			priority:    "",
			expectedErr: ErrInvalidPriority,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePriority(tt.priority)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ValidatePriority() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateStatus(t *testing.T) {
	tests := []struct {
		name        string
		status      string
		expectedErr error
	}{
		{
			name:        "Should accept PLANNING status",
			status:      "PLANNING",
			expectedErr: nil,
		},
		{
			name:        "Should accept PENDING status",
			status:      "PENDING",
			expectedErr: nil,
		},
		{
			name:        "Should accept PROGRESS status",
			status:      "IN_PROGRESS",
			expectedErr: nil,
		},
		{
			name:        "Should accept COMPLETED status",
			status:      "COMPLETED",
			expectedErr: nil,
		},
		{
			name:        "Should return Invalid status",
			status:      "Other",
			expectedErr: ErrInvalidStatus,
		}, {
			name:        "Should return Error Empty status",
			status:      "",
			expectedErr: ErrInvalidStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStatus(tt.status)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ValidateStatus() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}
