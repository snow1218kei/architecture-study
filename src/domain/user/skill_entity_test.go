package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
	tag "github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

func TestNewSkill(t *testing.T) {
	skillID := NewSkillID()
	createdAt := shared.NewCreatedAt()

	tests := []struct {
		testCase string
		params   SkillParams
		expected *Skill
		wantErr  error
	}{
		{
			testCase: "有効なparamsの場合",
			params: SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      2,
			},
			expected: &Skill{
				skillID:    skillID,
				tagID:      tag.TagID("1"),
				evaluation: 3,
				years:      2,
				createdAt:  createdAt,
			},
			wantErr: nil,
		},
		{
			testCase: "Evaluationが1未満の場合",
			params: SkillParams{
				TagID:      "1",
				Evaluation: 0,
				Years:      2,
			},
			expected: nil,
			wantErr:  errors.New("評価は1〜5の5段階です"),
		},
		{
			testCase: "Evaluationが5を超える場合",
			params: SkillParams{
				TagID:      "1",
				Evaluation: 6,
				Years:      2,
			},
			expected: nil,
			wantErr:  errors.New("評価は1〜5の5段階です"),
		},
		{
			testCase: "Yearsが1年未満の場合",
			params: SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      0,
			},
			expected: nil,
			wantErr:  errors.New("1年以上、5年以内で入力してください"),
		},
		{
			testCase: "Yearsが5を超える場合",
			params: SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      6,
			},
			expected: nil,
			wantErr:  errors.New("1年以上、5年以内で入力してください"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			skill, err := NewSkill(test.params, skillID, createdAt)

			assert.Equal(t, test.expected, skill)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
