package planusecase

type CreatePlanInput struct {
	Title              string
	Content            string
	Pricing            uint16
	Category           string
	UserID             string
	TagIDs             []string
	Status             string
	ConsultationMethod string
}
