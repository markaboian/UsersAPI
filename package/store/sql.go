package store

import (
	"UsersAPI/internal/domain"
	"database/sql"
)

type Sql struct {
	DB *sql.DB
}

func (s *Sql) GetUserById(id int) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE id = ?"

	row := s.DB.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.LastName, &user.PlaceOfBirth, &user.DateOfBirth, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}