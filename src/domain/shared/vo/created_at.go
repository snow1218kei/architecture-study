package shared

import (
	"errors"
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func NewCreatedAtByVal(val time.Time) (CreatedAt, error) {
	if val.IsZero() {
		return CreatedAt(time.Time{}), errors.New("createdAt must not be empty")
	}
	return CreatedAt(val), nil
}

func (createdAt CreatedAt) Value() time.Time {
	return time.Time(createdAt)
}
