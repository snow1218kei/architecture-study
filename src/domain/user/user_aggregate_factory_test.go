package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
)

func TestCreateUserAggregate(t *testing.T) {
	userAggregateFactory := UserAggregateFactory{
		UserParams: UserParams{
			Name:     "test",
			Email:    "test@example.com",
			Password: "abc123456def",
			Profile:  "test profile",
		},
		CareersParams: []CareerParams{
			{
				Detail:    "test career detail",
				StartYear: 2020,
				EndYear:   2022,
			},
		},
		SkillsParams: []SkillParams{
			{
				TagID:      "1",
				Evaluation: 3,
				Years:      2,
			},
		},
	}

	user, err := userAggregateFactory.CreateUserAggregate()

	assert.NoError(t, err)
	assert.NotEmpty(t, user.userID)
	assert.Equal(t, user.name, "test")
	assert.Equal(t, user.email, Email("test@example.com"))
	assert.Equal(t, user.password, Password("abc123456def"))
	assert.Equal(t, user.profile, "test profile")
	assert.Len(t, user.careers, 1)
	assert.Equal(t, user.careers[0].detail, "test career detail")
	assert.Equal(t, user.careers[0].startYear, uint16(2020))
	assert.Equal(t, user.careers[0].endYear, uint16(2022))
	assert.Len(t, user.skills, 1)
	assert.Equal(t, user.skills[0].tagID, tag.TagID("1"))
	assert.Equal(t, user.skills[0].evaluation, uint16(3))
	assert.Equal(t, user.skills[0].years, uint16(2))
}
