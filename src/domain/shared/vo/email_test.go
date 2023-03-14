package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		testCase string
		input    string
		expected shared.Email
		wantErr  string
	}{
		{
			testCase: "正常系：有効なメールアドレスの場合",
			input:    "test123@gmail.com",
			expected: shared.Email("test123@gmail.com"),
			wantErr:  "",
		},
		{
			testCase: "異常系：@が抜けている場合",
			input:    "test123gmail.com",
			expected: shared.Email(""),
			wantErr:  "無効なEmailアドレスです",
		},
		{
			testCase: "異常系：アドレスの先頭に.がある場合",
			input:    ".test123@gmail.com",
			expected: shared.Email(""),
			wantErr:  "無効なEmailアドレスです",
		},
		{
			testCase: "異常系：空文字の場合",
			input:    "",
			expected: shared.Email(""),
			wantErr:  "無効なEmailアドレスです",
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			actual, err := shared.NewEmail(test.input)
			assert.Equal(t, test.expected, actual)
			if test.wantErr == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, test.wantErr)
			}
		})
	}
}
