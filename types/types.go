package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*Users, error)
	GetUserByID(id int) (*Users, error)
	CreateUser(Users) error
}

type mockUserStore struct {
}

type Users struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, min=3, max=40"`
}

func GetUserByEmail(email string) (*Users, error) {
	return nil, nil
}
