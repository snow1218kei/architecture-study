package mentoringplan

import (
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"unicode/utf8"
)

const (
	PlanTitleMaxLength   = 255
	PlanContentMaxLength = 2000
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
	Category           string
	UserID             user.UserID
	TagIDs             []tag.TagID
	Status             string
	ConsultationMethod string
}

func NewMentoringPlan(params *MentoringPlanParams) (*MentoringPlan, error) {
	if err := validatePlanTitle(params.Title); err != nil {
		return nil, err
	}

	if err := validatePlanContent(params.Content); err != nil {
		return nil, err
	}

	category, err := shared.NewCategory(params.Category)
	if err != nil {
		return nil, err
	}

	consultationMethod, err := shared.NewConsultationMethod(params.ConsultationMethod)
	if err != nil {
		return nil, err
	}

	status, err := shared.NewStatus(params.Status)
	if err != nil {
		return nil, err
	}

	return &MentoringPlan{
		mentoringPlanID:    newMentoringPlanID(),
		userID:             params.UserID,
		title:              params.Title,
		content:            params.Content,
		pricing:            params.Pricing,
		category:           category,
		tagIDs:             params.TagIDs,
		status:             status,
		consultationMethod: consultationMethod,
	}, nil
}

func validatePlanTitle(title string) error {
	if utf8.RuneCountInString(title) > PlanTitleMaxLength {
		return apperr.BadRequestf("Messageが500文字を超えています: %s", title)
	}

	return nil
}

func validatePlanContent(content string) error {
	if utf8.RuneCountInString(content) > PlanContentMaxLength {
		return apperr.BadRequestf("Messageが500文字を超えています: %s", content)
	}

	return nil
}
