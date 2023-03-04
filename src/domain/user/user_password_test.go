package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		testCase string
		password string
		expected Password
		wantErr  string
	}{
		{
			testCase: "有効な入力値の場合",
			password: "abc123456def",
			expected: Password("abc123456def"),
			wantErr:  "",
		},
		{
			testCase: "入力値が12文字以下の場合",
			password: "abcde",
			expected: Password(""),
			wantErr:  "文字数は最低12文字以上でなければなりません",
		},
		{
			testCase: "英字が1文字も含まれていない場合",
			password: "123456789012333",
			expected: Password(""),
			wantErr:  "英字が最低1文字は含まれていなければなりません",
		},
		{
			testCase: "数字が1文字も含まれていない場合",
			password: "abcdefghijkjjj",
			expected: Password(""),
			wantErr:  "数字が最低1文字は含まれていなければなりません",
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			got, err := NewPassword(test.password)
			if test.wantErr != "" {
				assert.Empty(t, got)
				assert.EqualError(t, err, test.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}
