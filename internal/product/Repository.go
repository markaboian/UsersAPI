package product

import (
	"UsersAPI/internal/domain"
	"UsersAPI/package/store"
)

type IRepository interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error
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

func (r *Repository) UpdateUser(id int, user *domain.User) (*domain.User, error) {
	user, err := r.Interface.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(id int) error {
	err := r.Interface.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}