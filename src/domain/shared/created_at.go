package shared

import (
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (createdAt CreatedAt) Value() time.Time {
	return time.Time(createdAt)
}
