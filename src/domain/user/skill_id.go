package user

import (
	"github.com/google/uuid"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type SkillID string

func newSkillID() SkillID {
	return SkillID(uuid.New().String())
}

func NewSkillIDByVal(val string) (SkillID, error) {
	if val == "" {
		return SkillID(""), apperr.BadRequest("skillID must not be empty")
	}
	return SkillID(val), nil
}

func (skillId SkillID) String() string {
	return string(skillId)
}

func (skillId1 SkillID) Equal(skillId2 SkillID) bool {
	return skillId1 == skillId2
}
