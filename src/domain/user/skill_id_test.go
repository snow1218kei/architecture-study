package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewSkillID(t *testing.T) {
	tests := []struct {
		name string
		skillID user.SkillID
	}{
		{
			name:   "NewSkillID + String",
			skillID: user.NewSkillID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := uuid.Parse(tt.skillID.String())
			assert.Nil(t, err, "Expected no error, but got an error")
		})
	}
}

func TestNewSkillIDByVal(t *testing.T) {
	testCases := []struct {
		name      string
		val       string
		expected  user.SkillID
		expectErr bool
	}{
		{
			name:      "有効な値の場合",
			val:       "123",
			expected:  user.SkillID("123"),
			expectErr: false,
		},
		{
			name:      "空の値の場合",
			val:       "",
			expected:  user.SkillID(""),
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			skillID, err := user.NewSkillIDByVal(tt.val)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, skillID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, skillID)
			}
		})
	}
}

func TestSkillID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		skillID user.SkillID
	}{
		{
			name:   "Equal",
			skillID: user.NewSkillID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userID1 := tt.skillID
			userID2 := userID1
			assert.True(t, userID1.Equal(userID2), "Expected skillID1 and skillID2 to be equal, but they were not")
		})
	}
}
