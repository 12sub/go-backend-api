package user

import (
	"database/sql"
	"fmt"

	"github.com/go-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.Users, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.Users)
	for rows.Next() {
		_, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == "0" {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}
func scanRowIntoUser(rows *sql.Rows) (*types.Users, error) {
	users := new(types.Users)
	err := rows.Scan(
		&users.ID,
		&users.FirstName,
		&users.LastName,
		&users.Email,
		&users.Password,
		&users.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Store) GetUserByID(id int) (*types.Users, error) {
	return nil, nil
}

func (s *Store) CreateUser(users types.Users) error {
	return nil
}
