package requirement

import (
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type MentorRequirement struct {
	mentorID           MentorID
	title              string
	category           Category
	contractType       ContractType
	consultationMethod ConsultationMethod
	description        string
	budget             Budget
	applicationPeriod  ApplicationPeriod
	status             Status
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
	TagIDs             []uint16
	UserID             uint16
}

func NewMentorRequirement(params *MentorRequirementParams) (*MentorRequirement, error) {
	if err := validateTitle(params.Title); err != nil {
		return nil, err
	}

	if err := validateDescription(params.Description); err != nil {
		return nil, err
	}

	if err := validateCategory(Category(params.Category)); err != nil {
		return nil, err
	}

	if err := validateContractType(ContractType(params.ContractType)); err != nil {
		return nil, err
	}

	if err := validateStatus(Status(params.Status)); err != nil {
		return nil, err
	}

	budget, err := newBudget(&params.Budget)
	if err != nil {
		return nil, err
	}

	mentorReq := &MentorRequirement{
		mentorID:           newMentorID(),
		title:              params.Title,
		category:           Category(params.Category),
		contractType:       ContractType(params.ContractType),
		consultationMethod: ConsultationMethod(params.ConsultationMethod),
		description:        params.Description,
		budget:             *budget,
		applicationPeriod:  ApplicationPeriod(params.ApplicationPeriod),
		status:             Status(params.Status),
		tagIDs:             convertToTagIDs(params.TagIDs),
		userID:             user.UserID(params.UserID),
	}

	return mentorReq, nil
}

func validateTitle(title string) error {
	if len(title) == 0 || len(title) > 255 {
		return apperr.BadRequestf("Titleは0文字以上200文字以下である必要があります: %d", title)
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) == 0 || len(description) > 2000 {
		return apperr.BadRequestf("Descriptionは0文字以上2000文字以下である必要があります: %d", description)
	}

	return nil
}
