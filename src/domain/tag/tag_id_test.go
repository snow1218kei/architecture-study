package tag

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTagID(t *testing.T) {
	tests := []struct {
		name string
		tagID TagID
	}{
		{
			name:   "NewUserID + String",
			tagID: NewTagID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := uuid.Parse(tt.tagID.String())
			assert.Nil(t, err, "Expected no error, but got an error")
		})
	}
}

func TestTagID_Equal(t *testing.T) {
	tests := []struct {
		name    string
		tagID TagID
	}{
		{
			name:   "Equal",
			tagID: NewTagID(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tagID1 := tt.tagID
			tagID2 := tagID1
			assert.True(t, tagID1.Equal(tagID2), "Expected tagID1 and tagID2 to be equal, but they were not")
		})
	}
}
