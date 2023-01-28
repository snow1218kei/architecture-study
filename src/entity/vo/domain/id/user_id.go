type UserID string

func NewUserID() UserID {
	uuid.New()
}

func (userID UserID) String() string {
  string(userID)
}

func (userID1 UserID) Equal(userID2 UserID) bool {
  userID1 == userID2
}
