package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	t.Run("有効な入力値の場合, Password型の文字列を返す", func(t *testing.T) {
		expected := Password("abc123456def")
		actual, err := NewPassword("abc123456def")
		assert.Nil(t, err, "Expected no error, but got an error")
		assert.Equal(t, expected, actual, "Expected %s but got %s", expected, actual)
	})

	t.Run("入力値が12文字以下の場合, エラーを返す", func(t *testing.T) {
		_, err := NewPassword("abcde")
		assert.EqualError(t, err, "文字数は最低12文字以上でなければなりません")
	})

	t.Run("英字が1文字も含まれていない場合, エラーを返す", func(t *testing.T) {
		_, err := NewPassword("123456789012333")
		assert.EqualError(t, err, "英字が最低1文字は含まれていなければなりません")
	})

	t.Run("数字が1文字も含まれていない場合, エラーを返す", func(t *testing.T) {
		_, err := NewPassword("abcdefghijkjjj")
		assert.EqualError(t, err, "数字が最低1文字は含まれていなければなりません")
	})
}
