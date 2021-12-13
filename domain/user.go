package domain

import (
	"time"
)

type User struct {
	UserID     UserID
	FirstName  string
	MiddleName string
	BirthDay   time.Time
}

type UserID string

func NewUser(userId UserID, firstName string, middleName string, birth_day time.Time) (*User, error) {
	return &User{
		UserID:     userId,
		FirstName:  firstName,
		MiddleName: middleName,
		BirthDay:   birth_day,
	}, nil
}
