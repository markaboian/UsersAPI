package store

import "UsersAPI/internal/domain"

type Interface interface {
	GetUserById(id int) (*domain.User, error)
}