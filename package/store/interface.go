package store

import "UsersAPI/internal/domain"

type Interface interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
}