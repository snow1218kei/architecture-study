package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewUserID(t *testing.T) {
	tests := []struct {
		name string
		userID user.UserID
	}{
		{
			name:   "正常系：NewUserID + String",
			userID: user.NewUserID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := uuid.Parse(tt.userID.String())
			assert.Nil(t, err, "Expected no error, but got an error")
		})
	}
}

func TestNewUserIDByVal(t *testing.T) {
	testCases := []struct {
		name      string
		val       string
		expected  user.UserID
		expectErr bool
	}{
		{
			name:      "正常系：有効な値の場合",
			val:       "123",
			expected:  user.UserID("123"),
			expectErr: false,
		},
		{
			name:      "異常系：空の値の場合",
			val:       "",
			expected:  user.UserID(""),
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := user.NewUserIDByVal(tt.val)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, userID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, userID)
			}
		})
	}
}

func TestUserID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		userID user.UserID
	}{
		{
			name:   "正常系",
			userID: user.NewUserID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userID1 := tt.userID
			userID2 := userID1
			assert.True(t, userID1.Equal(userID2), "Expected userID1 and userID2 to be equal, but they were not")
		})
	}
}
