package user

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
)

func TestNewUser(t *testing.T) {
	userParams := UserParams{
		Name:     "test user",
		Email:    "test@example.com",
		Password: "password",
		Profile:  "test profile",
	}
	userID := NewUserID()
	email, _ := shared.NewEmail(userParams.Email)
	password, _ := NewPassword(userParams.Password)
	createdAt := shared.NewCreatedAt()
	careers := []*Career{{NewCareerID(), "companyA", 2015, 2020, createdAt}}
	skills := []*Skill{{NewSkillID(), "1", 4, 5, createdAt}}

	tests := []struct {
		testCase  string
		params    map[string]interface{}
		expected  *User
		wantError error
	}{
		{
			testCase: "有効なparamsの場合",
			params: map[string]interface{}{
				"userID":    userID,
				"name":      userParams.Name,
				"email":     email,
				"password":  password,
				"profile":   userParams.Profile,
				"careers":   careers,
				"skills":    skills,
				"createdAt": createdAt,
			},
			expected: &User{
				userID:    userID,
				name:      "test user",
				email:     email,
				password:  password,
				profile:   "test profile",
				careers:   careers,
				skills:    skills,
				createdAt: createdAt,
			},
			wantError: nil,
		},
		{
			testCase: "nameが長過ぎる場合",
			params: map[string]interface{}{
				"userID":    userID,
				"name":      strings.Repeat("s", 400),
				"email":     email,
				"password":  password,
				"profile":   "test profile",
				"careers":   careers,
				"skills":    skills,
				"createdAt": createdAt,
			},
			expected:  nil,
			wantError: errors.New("名前は255文字以下である必要があります。(現在400文字)"),
		},
		{
			testCase: "profileが長過ぎる場合",
			params: map[string]interface{}{
				"userID":    userID,
				"name":      "test user",
				"email":     email,
				"password":  password,
				"profile":   strings.Repeat("s", 3000),
				"careers":   careers,
				"skills":    skills,
				"createdAt": createdAt,
			},
			expected:  nil,
			wantError: errors.New("プロフィールは2000文字以下である必要があります。(現在3000文字)"),
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			user, err := NewUser(test.params)

			assert.Equal(t, test.wantError, err)
			assert.Equal(t, test.expected, user)
		})
	}
}
