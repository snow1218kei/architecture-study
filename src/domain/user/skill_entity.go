package user

import (
	"fmt"

	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	tag "github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

type Skill struct {
	skillID    SkillID
	tagID      tag.TagID
	evaluation uint16
	years      uint16
	createdAt  shared.CreatedAt
}

type SkillParams struct {
	TagID      string
	Evaluation uint16
	Years      uint16
}

func newSkill(params *SkillParams, skillID SkillID, createdAt shared.CreatedAt) (*Skill, error) {
	if err := validateEvalation(params.Evaluation); err != nil {
		return nil, err
	}

	if err := validateYears(params.Years); err != nil {
		return nil, err
	}

	tagID, err := tag.NewTagIDByVal(params.TagID)
	if err != nil {
			return nil, err
	}

	return &Skill{
		skillID:    skillID,
		tagID:      tagID,
		evaluation: params.Evaluation,
		years:      params.Years,
		createdAt:  createdAt,
	}, nil
}

func validateEvalation(evaluation uint16) error {
	if evaluation < 1 || 5 < evaluation {
		return fmt.Errorf("評価は1〜5の5段階です")
	}
	return nil
}

func validateYears(years uint16) error {
	if years < 1 || 5 < years {
		return fmt.Errorf("1年以上、5年以内で入力してください")
	}
	return nil
}
