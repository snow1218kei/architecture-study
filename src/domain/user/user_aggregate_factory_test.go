package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestCreateUserAggregate(t *testing.T) {
	tests := []struct {
		name          string
		userParams    user.UserParams
		careersParams []user.CareerParams
		skillsParams  []user.SkillParams
	}{
		{
			name: "正常系",
			userParams: user.UserParams{
				Name:     "test",
				Email:    "test@example.com",
				Password: "abc123456def",
				Profile:  "test profile",
			},
			careersParams: []user.CareerParams{
				{
					Detail:    "test career detail",
					StartYear: 2020,
					EndYear:   2022,
				},
			},
			skillsParams: []user.SkillParams{
				{
					TagID:      "1",
					Evaluation: 3,
					Years:      2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdUser, err := user.CreateUserAggregate(tt.userParams, tt.careersParams, tt.skillsParams)
			userForTest := user.GenFactoryForTest(tt.userParams, tt.careersParams, tt.skillsParams, createdUser)

			assert.NoError(t, err)
			assert.Equal(t, createdUser, userForTest)
		})
	}
}
