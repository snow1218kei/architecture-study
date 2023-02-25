package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
		t.Run("有効な入力値の場合,Email型のオブジェクトを返す", func(t *testing.T) {
				expected := Email("test123@gmail.com")
				actual, err := NewEmail("test123@gmail.com")
				assert.Nil(t, err)
				assert.Equal(t, expected, actual)
		})

		t.Run("ドメイン名にドットが含まれない場合、エラーを返す", func(t *testing.T) {
				_, err := NewEmail("test123@gmailcom")
				assert.NotNil(t, err)
		})

		t.Run("先頭にドットを含んだ場合、エラーを返す", func(t *testing.T) {
				_, err := NewEmail(".test123@gmailcom")
				assert.NotNil(t, err)
		})
}
