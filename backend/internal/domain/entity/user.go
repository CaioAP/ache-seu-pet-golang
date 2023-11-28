package entity

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Birthday  string
	Password  string
	CreatedAt time.Time
}

func NewUser(id *int, name string, email string, phone string, birthday string, createdAt time.Time) *User {
	return &User{
		ID:        *id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Birthday:  birthday,
		CreatedAt: createdAt,
	}
}
