package userinput

type UpdateUserInput struct {
	UserInput    UserIn      `json:"user"`
	CareersInput []*CareerIn `json:"career"`
	SkillsInput  []*SkillIn  `json:"skills"`
}

type UserIn struct {
	ID       string  `json:"userID"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Profile  *string `json:"profile"`
}

type CareerIn struct {
	ID        string  `json:"careerID"`
	Detail    *string `json:"detail"`
	StartYear *uint16 `json:"startYear"`
	EndYear   *uint16 `json:"endYear"`
}

type SkillIn struct {
	ID         string  `json:"skillID"`
	TagID      *string `json:"tagId"`
	Evaluation *uint16 `json:"evaluation"`
	Years      *uint16 `json:"year"`
}
