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
			name:        "Should not permit title with 3 characters or less.",
			title:       "abc",
			expectedErr: ErrInvalidTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTitle(tt.title)

			// 3. Verificamos se o erro retornado é igual ao esperado
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidateTitle() got = %v, expected %v", err, tt.expectedErr)
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

			// 3. Verificamos se o erro retornado é igual ao esperado
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidatePriority() got = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}
