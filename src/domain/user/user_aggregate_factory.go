package user

import "github.com/yuuki-tsujimura/architecture-study/src/domain/shared"

type UserAggregateFactory struct {
	UserParams    UserParams
	CareersParams []CareerParams
	SkillsParams  []SkillParams
}

func (factory UserAggregateFactory) CreateUserAggregate() (*User, error) {
	email, err := NewEmail(factory.UserParams.Email)
	password, err := NewPassword(factory.UserParams.Password)
	careers, err := prepareCareers(factory.CareersParams, shared.NewCreatedAt())
	skills, err := prepareSkills(factory.SkillsParams, shared.NewCreatedAt())
	userMap := map[string]interface{}{
		"userID":   NewUserID(),
		"name":     factory.UserParams.Name,
		"email":    email,
		"password": password,
		"profile":  factory.UserParams.Profile,
		"careers":  careers,
		"skills":   skills,
		"createdAt": shared.NewCreatedAt(),
	}
	user, err := NewUser(userMap)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func prepareCareers(careersParams []CareerParams, createdAt shared.CreatedAt) ([]*Career, error) {
	var careers []*Career
	for _, careerParams := range careersParams {
		career, err := NewCareer(careerParams, NewCareerID(), createdAt)
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
		skill, err := NewSkill(skillParams, NewSkillID(), createdAt)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}
