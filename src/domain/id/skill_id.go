package id

import (
	"github.com/google/uuid"
)

type SkillId string

func NewSkillId() SkillId {
	uuid, _ := uuid.NewRandom()
	return SkillId(uuid.String())
}

func (skillId SkillId) String() string {
	return string(skillId)
}

func (skillId1 SkillId) Equal(skillId2 SkillId) bool {
	return skillId1 == skillId2
}
