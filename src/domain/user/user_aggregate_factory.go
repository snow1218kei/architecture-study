package user

import (
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	tag "github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

func CreateUserAggregate(userParams UserParams, careersParams []CareerParams, skillsParams []SkillParams) (*User, error) {
	email, err := shared.NewEmail(userParams.Email)
	if err != nil {
		return nil, err
	}

	password, err := NewPassword(userParams.Password)
	if err != nil {
		return nil, err
	}

	careers, err := prepareCareers(careersParams, shared.NewCreatedAt())
	if err != nil {
		return nil, err
	}

	skills, err := prepareSkills(skillsParams, shared.NewCreatedAt())
	if err != nil {
		return nil, err
	}
	
	userInput := UserInput{
		Name:      userParams.Name,
		Email:     email,
		Password:  password,
		Profile:   userParams.Profile,
		Careers:   careers,
		Skills:    skills,
		CreatedAt: shared.NewCreatedAt(),
	}
	user, err := newUser(userInput)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func prepareCareers(careersParams []CareerParams, createdAt shared.CreatedAt) ([]*Career, error) {
	var careers []*Career
	for _, careerParams := range careersParams {
		career, err := newCareer(&careerParams, newCareerID(), createdAt)
		if err != nil {
			return nil, err
		}
		careers = append(careers, career)
	}
	return careers, nil
}

func prepareSkills(skillsParams []SkillParams, createdAt shared.CreatedAt) ([]*Skill, error) {
	var skills []*Skill
	for _, skillParams := range skillsParams {
		skill, err := newSkill(&skillParams, newSkillID(), createdAt)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

func GenCareerForTest(careerID CareerID, detail string, startYear uint16, endYear uint16, createdAt shared.CreatedAt) *Career {
	return &Career{
		careerID:  careerID,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
		createdAt: createdAt,
	}
}

func GenSkillForTest(skillID SkillID, tagID tag.TagID, evaluation uint16, years uint16, createdAt shared.CreatedAt) *Skill {
	return &Skill{
		skillID:    skillID,
		tagID:      tagID,
		evaluation: evaluation,
		years:      years,
		createdAt:  createdAt,
	}
}

func GenCareersForTest(careerID CareerID, detail string, startYear uint16, endYear uint16, createdAt shared.CreatedAt) []*Career {
	return []*Career{{
		careerID:  careerID,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
		createdAt: createdAt,
	}}
}

func GenSkillsForTest(skillID SkillID, tagID tag.TagID, evaluation uint16, years uint16, createdAt shared.CreatedAt) []*Skill {
	return []*Skill{{
		skillID:    skillID,
		tagID:      tagID,
		evaluation: evaluation,
		years:      years,
		createdAt:  createdAt,
	}}
}

func GenUserForTest(userID UserID, name string, email shared.Email, password Password, profile string, careers []*Career, skills []*Skill, createdAt shared.CreatedAt) *User {
	return &User{
		userID:    userID,
		name:      name,
		email:     email,
		password:  password,
		profile:   profile,
		careers:   careers,
		skills:    skills,
		createdAt: createdAt,
	}
}

func GenFactoryForTest(userParams UserParams, careerParams []CareerParams, skillParams []SkillParams, user *User) *User {
	careers := GenCareersForTest(user.careers[0].careerID, careerParams[0].Detail, careerParams[0].StartYear, careerParams[0].EndYear, user.careers[0].createdAt)
	skills := GenSkillsForTest(user.skills[0].skillID, tag.TagID(skillParams[0].TagID), skillParams[0].Evaluation, skillParams[0].Years, user.skills[0].createdAt)
	return GenUserForTest(user.userID, userParams.Name, shared.Email(userParams.Email), Password(userParams.Password), userParams.Profile, careers, skills, user.createdAt)
}
