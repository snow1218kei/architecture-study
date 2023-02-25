package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCareerID(t *testing.T) {
	t.Run("TestNewCareerID", func(t *testing.T) {
		careerID := NewCareerID()
		_, err := uuid.Parse(careerID.String())
		assert.Nil(t, err)
	})

	t.Run("TestString", func(t *testing.T) {
		careerID := NewCareerID()
		assert.True(t, reflect.TypeOf(careerID.String()).Kind() == reflect.String)
	})

	t.Run("TestEaqual", func(t *testing.T) {
		careerID1 := NewCareerID()
		careerID2 := careerID1
		assert.True(t, careerID1.Equal(careerID2))
	})
}
