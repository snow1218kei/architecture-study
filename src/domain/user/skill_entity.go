package user

import (
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/id"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase"
)

type Skill struct {
	SkillId    id.SkillId
	TagIds     []id.TagId
	Evaluation int
	Years      int
}

func NewSkill(input usecase.Skill) *Skill {
	skillId := id.NewSkillId()

	var tagIds []id.TagId
	for _, tagId := range input.TagIds {
		tagId, _ := id.GetTagId(tagId)
		tagIds = append(tagIds, tagId)
	}

	evaluation := input.Evaluation
	years := input.Years

	skill := &Skill{
		SkillId:    skillId,
		TagIds:     tagIds,
		Evaluation: evaluation,
		Years:      years,
	}

	skill.validateEvalation()
	skill.validateYears()

	return skill
}

func (skill Skill) validateEvalation() error {
	if skill.Evaluation < 1 || 5 < skill.Evaluation {
		return fmt.Errorf("評価は1〜5の5段階です")
	}

	return nil
}

func (skill Skill) validateYears() error {
	if skill.Years < 0 || 5 < skill.Years {
		return fmt.Errorf("0年以上、5年以内で入力してください")
	}

	return nil
}
