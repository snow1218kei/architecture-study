package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		testCase string
		input    string
		expected Email
		wantErr  string
	}{
		{
			testCase: "有効なメールアドレスの場合",
			input:    "test123@gmail.com",
			expected: Email("test123@gmail.com"),
			wantErr:  "",
		},
		{
			testCase: "@が抜けている場合",
			input:    "test123gmail.com",
			expected: Email(""),
			wantErr:  "無効なEmailアドレスです",
		},
		{
			testCase: "アドレスの先頭に.がある場合",
			input:    ".test123@gmail.com",
			expected: Email(""),
			wantErr:  "無効なEmailアドレスです",
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			actual, err := NewEmail(test.input)
			assert.Equal(t, test.expected, actual)
			if test.wantErr == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, test.wantErr)
			}
		})
	}
}
