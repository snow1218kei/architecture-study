package datamodel

import (
	"time"
)

type Career struct {
	ID        string       `db:"career_id"`
	UserID    int       `db:"user_id"`
	Detail    string    `db:"detail"`
	StartYear int       `db:"start_year"`
	EndYear   *int      `db:"end_year"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
