package planusecase

type GetPlanDTO struct {
	MentoringPlanID    string
	UserID             string
	Title              string
	Content            string
	pricing            uint16
	Category           string
	TagIDs             []string
	Status             string
	ConsultationMethod string
}
