package usecase

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type Career struct {
	Detail    string `json:"detail"`
	StartYear int    `json:"startYear"`
	EndYear   int    `json:"endYear"`
}

type Skill struct {
	TagIds     []int `json:"tagIds"`
	Evaluation int   `json:"evaluation"`
	Years      int   `json:"year"`
}

type CreateUserInput struct {
	User   User    `json:"user"`
	Career Career  `json:"career"`
	Skills []Skill `json:"skills"`
}
