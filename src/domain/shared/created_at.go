package shared

import (
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
  return CreatedAt(time.Now())
}

func (createdAt CreatedAt) String() string {
	return createdAt.String()
}

func (createdAt1 CreatedAt) Equal(createdAt2 CreatedAt) bool {
	return createdAt1 == createdAt2
}
