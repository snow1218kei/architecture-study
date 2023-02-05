package id

type TagId string

const (
	One = iota + 1
	Two
	Three
	Four
	Five
)

func GetTagId(id int) (TagId, bool) {
	switch id {
	case One:
		return TagId("1"), true
	case Two:
		return TagId("2"), true
	case Three:
		return TagId("3"), true
	case Four:
		return TagId("4"), true
	case Five:
		return TagId("5"), true
	default:
		return "", false
	}
}

func (tagId TagId) String() string {
	return string(tagId)
}

func (tagId1 TagId) Equal(tagId2 TagId) bool {
	return tagId1 == tagId2
}
