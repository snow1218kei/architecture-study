package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewCareerID(t *testing.T) {
	tests := []struct {
		name string
		careerID user.CareerID
	}{
		{
			name:   "正常系",
			careerID: user.NewCareerID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := uuid.Parse(tt.careerID.String())
			assert.Nil(t, err, "Expected no error, but got an error")
		})
	}
}

func TestNewCareerIDByVal(t *testing.T) {
	testCases := []struct {
		name      string
		val       string
		expected  user.CareerID
		expectErr bool
	}{
		{
			name:      "正常系：有効な値の場合",
			val:       "123",
			expected:  user.CareerID("123"),
			expectErr: false,
		},
		{
			name:      "異常系：空の値の場合",
			val:       "",
			expected:  user.CareerID(""),
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			careerID, err := user.NewCareerIDByVal(tt.val)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, careerID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, careerID)
			}
		})
	}
}

func TestCareerID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		careerID user.CareerID
	}{
		{
			name:   "正常系",
			careerID: user.NewCareerID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			careerID1 := tt.careerID
			careerID2 := careerID1
			assert.True(t, careerID1.Equal(careerID2), "Expected careerID1 and careerID2 to be equal, but they were not")
		})
	}
}
