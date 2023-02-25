package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func CreateUser() {
	inputJSON := `{
		"user": {
			"name": "John Doe",
			"email": "johndoe@example.com",
			"password": "secret",
			"profile": "Software Engineer"
		},
		"careers": [
			{
				"detail": "Working as a software engineer at XYZ Inc.",
				"startYear": 2018,
				"endYear": 2022
			},
			{
				"detail": "Worked as a software developer at ABC Ltd.",
				"startYear": 2015,
				"endYear": 2018
			}
		],
		"skills": [
			{
				"tagId": 1,
				"evaluation": 4,
				"year": 5
			},
			{
				"tagId": 2,
				"evaluation": 3,
				"year": 3
			}
		]
	}`

	var input CreateUserInput
	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		fmt.Println(err)
		return
	}

	userParams := user.UserParams{Name: input.UserInput.Name, Email: input.UserInput.Email, Password: input.UserInput.Password, Profile: input.UserInput.Profile}

	var careersParams []user.CareerParams
	for _, careerInput := range input.CareersInput {

		careerParams :=	user.CareerParams{Detail: careerInput.Detail, StartYear: careerInput.StartYear, EndYear: careerInput.EndYear}
			careersParams = append(careersParams, careerParams)
	}

	var skillsParams []user.SkillParams
	for _, skillInput := range input.SkillsInput {

		skillParams :=	user.SkillParams{TagID: skillInput.TagID, Evaluation: skillInput.Evaluation, Years: skillInput.Years}
			skillsParams = append(skillsParams, skillParams)
	}

	userAggregateFactory := user.UserAggregateFactory{
		UserParams:     userParams,
		CareersParams:  careersParams,
		SkillsParams:   skillsParams,
	}

	userAggregateFactory.CreateUserAggregate()
}
