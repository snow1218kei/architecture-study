package id

import (
	"github.com/google/uuid"
)

type CareerId string

func NewCareerId() CareerId {
	uuid, _ := uuid.NewRandom()
	return CareerId(uuid.String())
}

func (careerId CareerId) String() string {
	return string(careerId)
}

func (careerId1 CareerId) Equal(careerId2 CareerId) bool {
	return careerId1 == careerId2
}
