package product

import (
	"UsersAPI/internal/domain"
)

type IService interface {
	GetUserById(id int) (*domain.User, error)
	AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetUserById(id int) (*domain.User, error) {
	user, err := s.Repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) AddUser(name string, lastName string, placeOfBirth string, dateOfBirth string, email string) (*domain.User, error) {
	user, err := s.Repository.AddUser(name, lastName, placeOfBirth, dateOfBirth, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}