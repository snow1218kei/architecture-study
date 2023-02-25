package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSkillID(t *testing.T) {
	t.Run("TestNewSkillID", func(t *testing.T) {
		skillID := NewSkillID()
		_, err := uuid.Parse(skillID.String())
		assert.Nil(t, err)
	})

	t.Run("TestString", func(t *testing.T) {
		skillID := NewSkillID()
		assert.True(t, reflect.TypeOf(skillID.String()).Kind() == reflect.String)
	})

	t.Run("TestEaqual", func(t *testing.T) {
		skillID1 := NewSkillID()
		skillID2 := skillID1
		assert.True(t, skillID1.Equal(skillID2))
	})
}
