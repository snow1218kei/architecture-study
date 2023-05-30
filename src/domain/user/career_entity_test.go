package user_test

import (
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewCareer(t *testing.T) {
	careerID := user.NewCareerID()
	createdAt := shared.NewCreatedAt()

	tests := []struct {
		testCase string
		params   *user.CareerParams
		expected *user.Career
		wantErr  error
	}{
		{
			testCase: "正常系：有効なparamsの場合, career構造体のオブジェクトを返す",
			params: &user.CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2015,
				EndYear:   2022,
			},
			expected: user.GenCareerForTest(careerID, "Software Engineer", 2015, 2022, createdAt),
			wantErr:  nil,
		},
		{
			testCase: "異常系：Detailが長過ぎる場合, エラーを返す",
			params: &user.CareerParams{
				Detail:    strings.Repeat("s", 256),
				StartYear: 2015,
				EndYear:   2022,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("detailは255文字以下である必要があります。(現在256文字)"),
		},
		{
			testCase: "異常系：StartYearが1970年未満の場合, エラーを返す",
			params: &user.CareerParams{
				Detail:    "Software Engineer",
				StartYear: 1969,
				EndYear:   2022,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("startYearは1970年以上である必要があります: 1969"),
		},
		{
			testCase: "異常系：EndYearが1970年未満の場合, エラーを返す",
			params: &user.CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2015,
				EndYear:   1969,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("endYearは1970年以上であり、startYearより後の数値である必要があります: 1969"),
		},
		{
			testCase: "異常系：EndYearがStartYear以下の場合, エラーを返す",
			params: &user.CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2022,
				EndYear:   2015,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("endYearは1970年以上であり、startYearより後の数値である必要があります: 2015"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			career, err := user.NewCareer(test.params, careerID, createdAt)

			assert.Equal(t, test.expected, career)
			if err != nil && test.wantErr != nil {
				assert.Equal(t, test.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, test.wantErr, err)
			}
		})
	}
}
