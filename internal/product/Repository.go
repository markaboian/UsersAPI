package product

import (
	"UsersAPI/internal/domain"
	"UsersAPI/package/store"
)

type IRepository interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
}

type Repository struct {
	Interface store.Interface
}

func (r *Repository) GetUserById(id int) (*domain.User, error) {
	user, err := r.Interface.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error){
	user, err := r.Interface.AddUser(name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}