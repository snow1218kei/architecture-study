package plan

import (
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

const (
	MaxLength = 500
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
	subscription       *PlanSubscription
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
	if err := shared.ValidateCategory(params.Category); err != nil {
		return nil, err
	}

	if err := shared.ValidateConsultationMethod(params.ConsultationMethod); err != nil {
		return nil, err
	}

	if err := shared.ValidateStatus(params.Status); err != nil {
		return nil, err
	}

	return &MentoringPlan{
		mentoringPlanID:    newMentoringPlanID(),
		userID:             params.UserID,
		title:              params.Title,
		content:            params.Content,
		pricing:            params.Pricing,
		category:           shared.Category(params.Category),
		tagIDs:             params.TagIDs,
		status:             shared.Status(params.Status),
		consultationMethod: shared.ConsultationMethod(params.ConsultationMethod),
	}, nil
}

func (p *MentoringPlan) AddSubscriptionRequest(message string) error {
	subscriptionRequest, err := newSubscriptionRequest(message)
	if err != nil {
		return err
	}

	planSubscription, err := newPlanSubscription(*subscriptionRequest)
	if err != nil {
		return err
	}

	p.subscription = planSubscription

	return nil
}

func (p *MentoringPlan) AddSubscriptionApproval(message string) error {
	if p.subscription == nil {
		return apperr.Internal("subscription is not initialized")
	}

	subscriptionApproval, err := newSubscriptionApproval(message)
	if err != nil {
		return err
	}

	p.subscription.setSubscriptionApproval(subscriptionApproval)

	return nil
}

func (p *MentoringPlan) GetSubscription() *PlanSubscription {
	return p.subscription
}
