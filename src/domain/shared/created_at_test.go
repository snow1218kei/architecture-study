package shared

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatedAt(t *testing.T) {
		t.Run("TestNewCreatedAt", func(t *testing.T) {
			createdAt := NewCreatedAt()

			// 現在時刻から作成されたCreatedAtオブジェクトが返されたことを確認する
			assert.Equal(t, time.Now().UTC().Truncate(time.Second), time.Time(createdAt).UTC().Truncate(time.Second), "expected a new CreatedAt object to be created from the current time")
		})

		// t.Run("TestString", func(t *testing.T) {
		// 		createdAt := NewCreatedAt()
		// 		assert.True(t, reflect.TypeOf(createdAt.String()).Kind() == reflect.String)
		// })

		t.Run("TestEaqual", func(t *testing.T) {
				createdAtID1 := NewCreatedAt()
				createdAtID2 := createdAtID1
				assert.True(t, createdAtID1.Equal(createdAtID2))
		})
}
