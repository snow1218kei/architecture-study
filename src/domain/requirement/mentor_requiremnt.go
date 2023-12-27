package requirement

import (
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"unicode/utf8"
)

type MentorRequirement struct {
	mentorID           MentorID
	title              string
	category           shared.Category
	contractType       shared.ContractType
	consultationMethod shared.ConsultationMethod
	description        string
	budget             Budget
	applicationPeriod  ApplicationPeriod
	status             shared.Status
	tagIDs             []tag.TagID
	userID             user.UserID
}

type MentorRequirementParams struct {
	Title              string
	Category           string
	ContractType       string
	ConsultationMethod string
	Description        string
	Budget             BudgetParams
	ApplicationPeriod  string
	Status             string
	TagIDs             []tag.TagID
	UserID             user.UserID
}

func NewMentorRequirement(params *MentorRequirementParams) (*MentorRequirement, error) {
	if err := validateTitle(params.Title); err != nil {
		return nil, err
	}

	if err := validateDescription(params.Description); err != nil {
		return nil, err
	}

	category, err := shared.NewCategory(params.Category)
	if err != nil {
		return nil, err
	}

	contractType, err := shared.NewContractType(params.ContractType)
	if err != nil {
		return nil, err
	}

	status, err := shared.NewStatus(params.Status)
	if err != nil {
		return nil, err
	}

	consultationMethod, err := shared.NewConsultationMethod(params.ConsultationMethod)
	if err != nil {
		return nil, err
	}

	applicationPeriod, err := newApplicationPeriod(params.ApplicationPeriod)
	if err != nil {
		return nil, err
	}

	budget, err := newBudget(&params.Budget)
	if err != nil {
		return nil, err
	}

	mentorReq := &MentorRequirement{
		mentorID:           newMentorID(),
		title:              params.Title,
		category:           category,
		contractType:       contractType,
		consultationMethod: consultationMethod,
		description:        params.Description,
		budget:             *budget,
		applicationPeriod:  applicationPeriod,
		status:             status,
		tagIDs:             params.TagIDs,
		userID:             params.UserID,
	}

	return mentorReq, nil
}

func validateTitle(title string) error {
	if utf8.RuneCountInString(title) == 0 || utf8.RuneCountInString(title) > 255 {
		return apperr.BadRequestf("Titleは0文字以上200文字以下である必要があります: %s", title)
	}

	return nil
}

func validateDescription(description string) error {
	if utf8.RuneCountInString(description) == 0 || utf8.RuneCountInString(description) > 2000 {
		return apperr.BadRequestf("Descriptionは0文字以上2000文字以下である必要があります: %s", description)
	}

	return nil
}
