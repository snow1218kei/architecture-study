package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		testCase string
		password string
		expected user.Password
		wantErr  string
	}{
		{
			testCase: "正常系：有効な入力値の場合",
			password: "abc123456def",
			expected: user.Password("abc123456def"),
			wantErr:  "",
		},
		{
			testCase: "異常系：入力値が12文字以下の場合",
			password: "abcde",
			expected: user.Password(""),
			wantErr:  "文字数は最低12文字以上でなければなりません",
		},
		{
			testCase: "異常系：英字が1文字も含まれていない場合",
			password: "123456789012333",
			expected: user.Password(""),
			wantErr:  "英字が最低1文字は含まれていなければなりません",
		},
		{
			testCase: "異常系：数字が1文字も含まれていない場合",
			password: "abcdefghijkjjj",
			expected: user.Password(""),
			wantErr:  "数字が最低1文字は含まれていなければなりません",
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			got, err := user.NewPassword(test.password)
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
