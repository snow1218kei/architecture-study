package user

import (
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
	email, _ := NewEmail(userParams.Email)
	password, _ := NewPassword(userParams.Password)
	createdAt := shared.NewCreatedAt()
	careers := []*Career{{NewCareerID(), "companyA", 2015, 2020, createdAt}}
	skills := []*Skill{{NewSkillID(), "1", 4, 5, createdAt}}

	t.Run("有効なparamsの場合, user構造体のオブジェクトを返す", func(t *testing.T) {	
		userMap := map[string]interface{}{
			"userID":     userID,
			"name":       userParams.Name,
			"email":      email,
			"password":   password,
			"profile":    userParams.Profile,
			"careers":    careers,
			"skills":     skills,
			"createdAt":  createdAt,
		}
	
		user, err := NewUser(userMap)

		assert.Nil(t, err)
		assert.Equal(t, userID, user.userID)
		assert.Equal(t, "test user", user.name)
		assert.Equal(t, email, user.email)
		assert.Equal(t, password, user.password)
		assert.Equal(t, "test profile", user.profile)
		assert.Len(t, user.careers, 1)
		assert.Len(t, user.skills, 1)
	})

	t.Run("nameが長過ぎる場合, エラーを返す", func(t *testing.T) {
		userMap := map[string]interface{}{
			"userID":     userID,
			"name":       strings.Repeat("s", 400),
			"email":      email,
			"password":   password,
			"profile":    "test profile",
			"careers":    careers,
			"skills":     skills,
			"createdAt":  createdAt,
		}
		
		user, err := NewUser(userMap)

		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("profileが長過ぎる場合, エラーを返す", func(t *testing.T) {
		userMap := map[string]interface{}{
			"userID":     userID,
			"name":       "test user",
			"email":      email,
			"password":   password,
			"profile":    strings.Repeat("s", 3000),
			"careers":    careers,
			"skills":     skills,
			"createdAt":  createdAt,
		}
		
		user, err := NewUser(userMap)

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
