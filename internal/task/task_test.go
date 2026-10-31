package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	t.Run("Should create task with default status and filled date", func(t *testing.T) {
		title := "Title example"
		description := "Description example"
		priority := PriorityLow
		status := StatusPending

		_, err := NewTask(title, description, priority, status)
		if err != nil {
			t.Errorf("Expected: created task, got %v", err)
		}
	})

	t.Run("Should create task with no description", func(t *testing.T) {
		title := "Title example"
		description := ""
		priority := PriorityLow
		status := StatusPending

		_, err := NewTask(title, description, priority, status)
		if err != nil {
			t.Errorf("Expected: created task, got %v", err)
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
			title:       "abc",
			expectedErr: ErrTitleTooShort,
		},
		{
			name:        "Should not permit title with more then 50 characters",
			title:       "This is a long title for a task created and should not pass.",
			expectedErr: ErrTitleTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTitle(tt.title)

			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
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
			description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus bibendum sagittis est pretium fringilla. Nam ultricies blandit enim, sit amet aliquet purus vestibulum sed. In finibus eleifend sapien. End",
			expectedErr: ErrDescriptionTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDescription(tt.description)

			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidateDescription() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidatePriority(t *testing.T) {
	tests := []struct {
		name        string
		priority    Priority
		expectedErr error
	}{
		{
			name:        "Should accept LOW priority",
			priority:    PriorityLow,
			expectedErr: nil,
		},
		{
			name:        "Should accept MEDIUM priority",
			priority:    PriorityMedium,
			expectedErr: nil,
		},
		{
			name:        "Should accept HIGH priority",
			priority:    PriorityHigh,
			expectedErr: nil,
		},
		{
			name:        "Should return Invalid Priority",
			priority:    Priority("URGENT"),
			expectedErr: ErrInvalidPriority,
		}, {
			name:        "Should return Error Empty Priority",
			priority:    Priority(""),
			expectedErr: ErrInvalidPriority,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePriority(tt.priority)

			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidatePriority() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateStatus(t *testing.T) {
	tests := []struct {
		name        string
		status      Status
		expectedErr error
	}{
		{
			name:        "Should accept PLANNING status",
			status:      StatusPlanning,
			expectedErr: nil,
		},
		{
			name:        "Should accept PENDING status",
			status:      StatusPending,
			expectedErr: nil,
		},
		{
			name:        "Should accept PROGRESS status",
			status:      StatusProgress,
			expectedErr: nil,
		},
		{
			name:        "Should accept COMPLETED status",
			status:      StatusCompleted,
			expectedErr: nil,
		},
		{
			name:        "Should return Invalid status",
			status:      Status("Other"),
			expectedErr: ErrInvalidStatus,
		}, {
			name:        "Should return Error Empty status",
			status:      Status(""),
			expectedErr: ErrInvalidStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStatus(tt.status)

			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidateStatus() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}
