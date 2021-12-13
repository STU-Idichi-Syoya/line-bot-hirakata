package domain

type UserAttr struct {
	UserID UserID
	Hobby  string
}

func NewUserAttr(userId UserID, Hobby string) (*UserAttr, error) {
	return &UserAttr{
		UserID: userId,
		Hobby:  Hobby,
	}, nil
}
