package datamodel

import (
	"time"
)

type Skill struct {
	SkillID    int       `db:"skill_id"`
	UserID     int       `db:"user_id"`
	TagID      int       `db:"tag_id"`
	Evaluation int       `db:"evaluation"`
	Years      int       `db:"years"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
