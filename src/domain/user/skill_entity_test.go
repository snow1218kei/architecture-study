package user_test

import (
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"testing"

	"github.com/stretchr/testify/assert"
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	tag "github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewSkill(t *testing.T) {
	skillID := user.NewSkillID()
	createdAt := shared.NewCreatedAt()

	tests := []struct {
		testCase string
		params   *user.SkillParams
		expected *user.Skill
		wantErr  error
	}{
		{
			testCase: "正常系：有効なparamsの場合",
			params: &user.SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      2,
			},
			expected: user.GenSkillForTest(skillID, tag.TagID("1"), 3, 2, createdAt),
			wantErr:  nil,
		},
		{
			testCase: "異常系：Evaluationが1未満の場合",
			params: &user.SkillParams{
				TagID:      "1",
				Evaluation: 0,
				Years:      2,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("評価は1〜5の5段階です: 0"),
		},
		{
			testCase: "異常系：Evaluationが5を超える場合",
			params: &user.SkillParams{
				TagID:      "1",
				Evaluation: 6,
				Years:      2,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("評価は1〜5の5段階です: 6"),
		},
		{
			testCase: "異常系：Yearsが1年未満の場合",
			params: &user.SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      0,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("1年以上、5年以内で入力してください: 0"),
		},
		{
			testCase: "異常系：Yearsが5を超える場合",
			params: &user.SkillParams{
				TagID:      "1",
				Evaluation: 3,
				Years:      6,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("1年以上、5年以内で入力してください: 6"),
		},
		{
			testCase: "異常系：TagIDが空の場合",
			params: &user.SkillParams{
				TagID:      "",
				Evaluation: 3,
				Years:      2,
			},
			expected: nil,
			wantErr:  apperr.BadRequest("tagID must not be empty"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			skill, err := user.NewSkill(test.params, skillID, createdAt)

			assert.Equal(t, test.expected, skill)
			if err != nil && test.wantErr != nil {
				assert.Equal(t, test.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, test.wantErr, err)
			}
		})
	}
}
