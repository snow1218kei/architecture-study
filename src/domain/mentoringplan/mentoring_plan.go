package mentoringplan

import (
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"unicode/utf8"
)

const (
	planTitleMaxLength   = 255
	planContentMaxLength = 2000
)

type MentoringPlan struct {
	mentoringPlanID    MentoringPlanID
	userID             user.UserID
	title              string
	content            string
	pricing            uint16
	category           shared.Category
	tagIDs             []tag.TagID
	status             shared.Status
	consultationMethod shared.ConsultationMethod
}

type MentoringPlanParams struct {
	Title              string
	Content            string
	Pricing            uint16
	Category           shared.Category
	UserID             user.UserID
	TagIDs             []tag.TagID
	Status             shared.Status
	ConsultationMethod shared.ConsultationMethod
}

func NewMentoringPlan(params *MentoringPlanParams) (*MentoringPlan, error) {
	if err := validatePlanTitle(params.Title); err != nil {
		return nil, err
	}

	if err := validatePlanContent(params.Content); err != nil {
		return nil, err
	}

	return &MentoringPlan{
		mentoringPlanID:    newMentoringPlanID(),
		userID:             params.UserID,
		title:              params.Title,
		content:            params.Content,
		pricing:            params.Pricing,
		category:           params.Category,
		tagIDs:             params.TagIDs,
		status:             params.Status,
		consultationMethod: params.ConsultationMethod,
	}, nil
}

func validatePlanTitle(title string) error {
	if utf8.RuneCountInString(title) > planTitleMaxLength {
		return apperr.BadRequestf("Messageが%d文字を超えています: %d", planTitleMaxLength, title)
	}

	return nil
}

func validatePlanContent(content string) error {
	if utf8.RuneCountInString(content) > planContentMaxLength {
		return apperr.BadRequestf("Messageが%d文字を超えています: %d", planContentMaxLength, content)
	}

	return nil
}
