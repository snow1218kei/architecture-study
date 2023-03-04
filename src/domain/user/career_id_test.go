package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCareerID(t *testing.T) {
	tests := []struct {
		name string
		careerID CareerID
	}{
		{
			name:   "NewUserID + String",
			careerID: NewCareerID(),
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

func TestCareerID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		careerID CareerID
	}{
		{
			name:   "Equal",
			careerID: NewCareerID(),
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
