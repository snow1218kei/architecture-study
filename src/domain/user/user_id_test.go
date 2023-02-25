package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserID(t *testing.T) {
	t.Run("TestNewUserID", func(t *testing.T) {
		userID := NewUserID()
		_, err := uuid.Parse(userID.String())
		assert.Nil(t, err, "Expected no error, but got an error")
	})

	t.Run("TestString", func(t *testing.T) {
		userID := NewUserID()
		assert.True(t, reflect.TypeOf(userID).Kind() == reflect.String, "Expected type string, but got a different type")
	})

	t.Run("TestEqual", func(t *testing.T) {
		userID1 := NewUserID()
		userID2 := userID1
		assert.True(t, userID1.Equal(userID2), "Expected userID1 and userID2 to be equal, but they were not")
	})
}
