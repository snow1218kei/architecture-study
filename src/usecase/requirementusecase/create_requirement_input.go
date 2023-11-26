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
	TagIDs             []uint16
	UserID             uint16
}

type CreateBudgetInput struct {
	LowerBound uint16
	UpperBound uint16
}
