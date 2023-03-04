package shared

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCreatedAt(t *testing.T) {
	tests := []struct {
		name      string
		createdAt CreatedAt
	}{
		{
			name:      "NewCreatedAt",
			createdAt: NewCreatedAt(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			if now.Before(tt.createdAt.Value()) {
				t.Errorf("NewCreatedAt() returned invalid value: %v", tt.createdAt.Value())
			}
		})
	}
}

func TestCreatedAt_Value(t *testing.T) {
	tests := []struct {
		name      string
		createdAt CreatedAt
		expected  time.Time
	}{
		{
			name:      "Value()",
			createdAt: CreatedAt(time.Now()),
			expected:  time.Now(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.createdAt.Value())
		})
	}
}
