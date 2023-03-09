package user

import (
	"log"
)

type Service interface {
	Create(firstName, lastName, email, telefono string) (*User, error)
	GetAll(filers filter) ([]User, error)
	GetById(id string) (*User, error)
	Delete(id string) (bool, error)
	Update(id string, firstName *string, lastName *string, email *string, telefono *string) error
}

type (
	service struct {
		log  *log.Logger
		repo Repository
	}

	filter struct {
		First_name string
		Last_name  string
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, telefono string) (*User, error) {
	s.log.Println("Create user service")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Telefono:  telefono,
	}
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s service) GetAll(filters filter) ([]User, error) {
	var users []User
	users, err := s.repo.GetAll(filters)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) GetById(id string) (*User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) Delete(id string) (bool, error) {
	value, err := s.repo.Delete(id)
	if err != nil {
		return false, err
	}
	return value, nil
}

func (service service) Update(id string, firstName *string, lastName *string, email *string, telefono *string) error {
	return service.repo.Update(id, firstName, lastName, email, telefono)
}
