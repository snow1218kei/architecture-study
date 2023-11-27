package requirementusecase

type CreateRequirementInput struct {
	Title              string
	Category           string
	ContractType       string
	ConsultationMethod string
	Description        string
	Budget             CreateBudgetInput
	ApplicationPeriod  string
	Status             string
	TagIDs             []string
	UserID             string
}

type CreateBudgetInput struct {
	LowerBound uint16
	UpperBound uint16
}
