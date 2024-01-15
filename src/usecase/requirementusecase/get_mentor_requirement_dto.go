package requirementusecase

type GetMentorRequirementDTO struct {
	MentorID           uint16
	Title              string
	Category           string
	ContractType       string
	ConsultationMethod string
	Description        string
	BudgetLowerBound   uint16
	BudgetUpperBound   uint16
	ApplicationPeriod  string
	Status             string
	TagIDs             []string
	UserID             string
}
