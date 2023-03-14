package userinput

type CreateUserInput struct {
	UserInput   UserInput    `json:"user"`
	CareersInput []*CareerInput  `json:"career"`
	SkillsInput []*SkillInput `json:"skills"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type CareerInput struct {
	Detail    string `json:"detail"`
	StartYear uint16    `json:"startYear"`
	EndYear   uint16    `json:"endYear"`
}

type SkillInput struct {
	TagID      string `json:"tagId"`
	Evaluation uint16   `json:"evaluation"`
	Years      uint16   `json:"year"`
}
