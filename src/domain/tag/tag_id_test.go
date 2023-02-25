package tag

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTagID(t *testing.T) {
	t.Run("TestNewTagID", func(t *testing.T) {
		tagID := NewTagID()
		_, err := uuid.Parse(tagID.String())
		assert.Nil(t, err)
	})

	t.Run("TestString", func(t *testing.T) {
		tagID := NewTagID()
		assert.True(t, reflect.TypeOf(tagID).Kind() == reflect.String)
	})

	t.Run("TestEaqual", func(t *testing.T) {
		tagID1 := NewTagID()
		tagID2 := tagID1
		assert.True(t, tagID1.Equal(tagID2))
	})
}
