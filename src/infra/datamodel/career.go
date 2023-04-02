package datamodel

import (
	"time"
)

type Career struct {
	ID        int       `db:"career_id" json:"career_id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Detail    string    `db:"detail" json:"detail"`
	StartYear int       `db:"start_year" json:"start_year"`
	EndYear   *int      `db:"end_year" json:"end_year"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
