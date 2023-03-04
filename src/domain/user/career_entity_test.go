package user

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
)

func TestNewCareer(t *testing.T) {
	careerID := NewCareerID()
	createdAt := shared.NewCreatedAt()

	tests := []struct {
		testCase string
		params   CareerParams
		expected *Career
		wantErr  error
	}{
		{
			testCase: "有効なparamsの場合, career構造体のオブジェクトを返す",
			params: CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2015,
				EndYear:   2022,
			},
			expected: &Career{
				careerID:  careerID,
				detail:    "Software Engineer",
				startYear: 2015,
				endYear:   2022,
				createdAt: createdAt,
			},
			wantErr: nil,
		},
		{
			testCase: "Detailが長過ぎる場合, エラーを返す",
			params: CareerParams{
				Detail:    strings.Repeat("s", 256),
				StartYear: 2015,
				EndYear:   2022,
			},
			expected: nil,
			wantErr:  errors.New("detailは255文字以下である必要があります。(現在256文字)"),
		},
		{
			testCase: "StartYearが1970年未満の場合, エラーを返す",
			params: CareerParams{
				Detail:    "Software Engineer",
				StartYear: 1969,
				EndYear:   2022,
			},
			expected: nil,
			wantErr:  errors.New("startYearは1970年以上である必要があります"),
		},
		{
			testCase: "EndYearが1970年未満の場合, エラーを返す",
			params: CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2015,
				EndYear:   1969,
			},
			expected: nil,
			wantErr:  errors.New("endYearは1970年以上であり、startYearより後の数値である必要があります"),
		},
		{
			testCase: "EndYearがStartYear以下の場合, エラーを返す",
			params: CareerParams{
				Detail:    "Software Engineer",
				StartYear: 2022,
				EndYear:   2015,
			},
			expected: nil,
			wantErr:  errors.New("endYearは1970年以上であり、startYearより後の数値である必要があります"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			career, err := NewCareer(test.params, careerID, createdAt)

			assert.Equal(t, test.expected, career)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
