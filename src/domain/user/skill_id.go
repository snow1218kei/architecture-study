package user

import (
	"errors"

	"github.com/google/uuid"
)

type SkillID string

func newSkillID() SkillID {
	return SkillID(uuid.New().String())
}

func NewSkillIDByVal(val string) (SkillID, error) {
	if val == "" {
		return SkillID(""), errors.New("skillID must not be empty")
	}
	return SkillID(val), nil
}

func (skillId SkillID) String() string {
	return string(skillId)
}

func (skillId1 SkillID) Equal(skillId2 SkillID) bool {
	return skillId1 == skillId2
}
