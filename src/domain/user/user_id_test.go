package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserID(t *testing.T) {
	tests := []struct {
		name string
		userID UserID
	}{
		{
			name:   "NewUserID + String",
			userID: NewUserID(),
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

func TestUserID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		userID UserID
	}{
		{
			name:   "Equal",
			userID: NewUserID(),
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
