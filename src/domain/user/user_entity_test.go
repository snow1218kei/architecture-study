package user_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
)

func TestNewUser(t *testing.T) {
	userParams := user.UserParams{
		Name:     "test user",
		Email:    "test@example.com",
		Password: "password",
		Profile:  "test profile",
	}
	userID := user.NewUserID()
	email, _ := shared.NewEmail(userParams.Email)
	password, _ := user.NewPassword(userParams.Password)
	createdAt := shared.NewCreatedAt()
	careers := user.GenCareersForTest(user.NewCareerID(), "companyB", 2010, 2015, createdAt)
	skills := user.GenSkillsForTest(user.NewSkillID(), "1", 4, 5, createdAt)
	tests := []struct {
		testCase  string
		input    user.UserInput
		expected  *user.User
		wantError error
	}{
		{
			testCase: "有効なparamsの場合",
			input: user.UserInput{
				UserID:    userID,
				Name:      userParams.Name,
				Email:     email,
				Password:  password,
				Profile:   userParams.Profile,
				Careers:   careers,
				Skills:    skills,
				CreatedAt: createdAt,
			},
			expected: user.GenUserForTest(userID, "test user", email, password, "test profile", careers, skills, createdAt),
			wantError: nil,
		},
		{
			testCase: "nameが長過ぎる場合",
			input: user.UserInput{
				UserID:    userID,
				Name:      strings.Repeat("s", 400),
				Email:     email,
				Password:  password,
				Profile:   "test profile",
				Careers:   careers,
				Skills:    skills,
				CreatedAt: createdAt,
			},
			expected:  nil,
			wantError: errors.New("名前は255文字以下である必要があります。(現在400文字)"),
		},
		{
			testCase: "profileが長過ぎる場合",
			input: user.UserInput{
				UserID:    userID,
				Name:      "test user",
				Email:     email,
				Password:  password,
				Profile:   strings.Repeat("s", 3000),
				Careers:   careers,
				Skills:    skills,
				CreatedAt: createdAt,
			},
			expected:  nil,
			wantError: errors.New("プロフィールは2000文字以下である必要があります。(現在3000文字)"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			user, err := user.NewUser(test.input)

			assert.Equal(t, test.wantError, err)
			assert.Equal(t, test.expected, user)
		})
	}
}
