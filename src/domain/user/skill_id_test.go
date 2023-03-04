package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillID(t *testing.T) {
	tests := []struct {
		name string
		skillID SkillID
	}{
		{
			name:   "NewSkillID + String",
			skillID: NewSkillID(),
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

func TestSkillID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		skillID SkillID
	}{
		{
			name:   "Equal",
			skillID: NewSkillID(),
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
