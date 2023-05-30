package shared

import (
	"time"

	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func NewCreatedAtByVal(val time.Time) (CreatedAt, error) {
	if val.IsZero() {
		return CreatedAt(time.Time{}), apperr.BadRequest("createdAt must not be empty")
	}
	return CreatedAt(val), nil
}

func (createdAt CreatedAt) Value() time.Time {
	return time.Time(createdAt)
}
