package datamodel

import (
	"time"
)

type Skill struct {
	ID    string       `db:"skill_id"`
	UserID     string       `db:"user_id"`
	TagID      string       `db:"tag_id"`
	Evaluation int       `db:"evaluation"`
	Years      int       `db:"years"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
