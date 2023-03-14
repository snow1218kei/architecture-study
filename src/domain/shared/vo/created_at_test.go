package shared_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
)

func TestNewCreatedAt(t *testing.T) {
	tests := []struct {
		name      string
		createdAt shared.CreatedAt
	}{
		{
			name:      "正常系",
			createdAt: shared.NewCreatedAt(),
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

func TestNewCreatedAByVal(t *testing.T) {
	testCases := []struct {
		name      string
		val       time.Time
		expected  shared.CreatedAt
		expectErr bool
	}{
		{
			name:      "有効な値の場合",
			val:       time.Date(2023, 3, 9, 12, 30, 0, 0, time.Local),
			expected:  shared.CreatedAt(time.Date(2023, 3, 9, 12, 30, 0, 0, time.Local)),
			expectErr: false,
		},
		{
			name:      "ゼロ値の場合",
			val:       time.Time{},
			expected:  shared.CreatedAt(time.Time{}),
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			createdAt, err := shared.NewCreatedAtByVal(tt.val)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, createdAt)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, createdAt)
			}
		})
	}
}
