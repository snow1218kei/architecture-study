package tag

import (
	"errors"

	"github.com/google/uuid"
)

type TagID string

func NewTagID() TagID {
	return TagID(uuid.New().String())
}

func NewTagIDByVal(val string) (TagID, error) {
	if val == "" {
		return TagID(""), errors.New("tagID must not be empty")
	}
	return TagID(val), nil
}

func (tagId TagID) String() string {
	return string(tagId)
}

func (tagId1 TagID) Equal(tagId2 TagID) bool {
	return tagId1 == tagId2
}
