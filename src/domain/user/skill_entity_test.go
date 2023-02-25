package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
	tag "github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

func TestNewSkill(t *testing.T) {
	skillID := NewSkillID()
	createdAt := shared.NewCreatedAt()
	
	t.Run("有効なparamsの場合, skill構造体のオブジェクトを返す", func(t *testing.T) {
		params := SkillParams{
			TagID:      "1",
			Evaluation: 3,
			Years:      2,
		}

		skill, err := NewSkill(params, skillID, createdAt)

		assert.NoError(t, err)
		assert.Equal(t, skillID, skill.skillID)
		assert.Equal(t, tag.TagID("1"), skill.tagID)
		assert.Equal(t, uint16(3), skill.evaluation)
		assert.Equal(t, uint16(2), skill.years)
		assert.Equal(t, createdAt, skill.createdAt)
	})

	t.Run("Evauationが1未満の場合、エラーを返す", func(t *testing.T) {
		params := SkillParams{
			TagID:      "1",
			Evaluation: 0,
			Years:      2,
		}

		skill, err := NewSkill(params, skillID, createdAt)
		assert.Error(t, err)
		assert.Nil(t, skill)
	})

	t.Run("Evaluationが5を超える場合、エラーを返す", func(t *testing.T) {
		params := SkillParams{
			TagID:      "1",
			Evaluation: 6,
			Years:      2,
		}

		skill, err := NewSkill(params, skillID, createdAt)
		assert.Error(t, err)
		assert.Nil(t, skill)
	})

	t.Run("Yearsが1年未満の場合、エラーを返す", func(t *testing.T) {
		params := SkillParams{
			TagID:      "1",
			Evaluation: 3,
			Years:      0,
		}

		skill, err := NewSkill(params, skillID, createdAt)
		assert.Error(t, err)
		assert.Nil(t, skill)
	})

	t.Run("経験年数が5を超える場合、エラーを返す", func(t *testing.T) {
		params := SkillParams{
			TagID:      "1",
			Evaluation: 3,
			Years:      6,
		}

	  skill, err := NewSkill(params, skillID, createdAt)
		assert.Error(t, err)
		assert.Nil(t, skill)
	})
}
