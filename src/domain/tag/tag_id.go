package tag

import (
	"github.com/google/uuid"
)

type TagID string

func NewTagID() TagID {
	return TagID(uuid.New().String())
}

func (tagId TagID) String() string {
	return string(tagId)
}

func (tagId1 TagID) Equal(tagId2 TagID) bool {
	return tagId1 == tagId2
}
