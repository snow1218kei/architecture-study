package datamodel

import (
	"time"
)

type Skill struct {
	SkillID    int       `db:"skill_id" json:"skill_id"`
	UserID     int       `db:"user_id" json:"user_id"`
	TagID      int       `db:"tag_id" json:"tag_id"`
	Evaluation int       `db:"evaluation" json:"evaluation"`
	Years      int       `db:"years" json:"years"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
