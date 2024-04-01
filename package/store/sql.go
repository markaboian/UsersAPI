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

func(s *Sql) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error) {

	query := "INSERT INTO users (name, lastName, place_of_birth, date_of_birth, email) VALUES (?, ?, ?, ?, ?)"

	result, err := s.DB.Exec(query, name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Id: int(id),
		Name: name,
		LastName: lastName,
		PlaceOfBirth: placeOfBirth,
		DateOfBirth: dateOfBirth,
		Email: email,
	}

	return user, nil
}