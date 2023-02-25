package user

import (
	"github.com/google/uuid"
)

type SkillID string

func NewSkillID() SkillID {
  return SkillID(uuid.New().String())
}

func (skillId SkillID) String() string {
	return string(skillId)
}

func (skillId1 SkillID) Equal(skillId2 SkillID) bool {
	return skillId1 == skillId2
}
