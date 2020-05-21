package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserService interface {
	CreateUser(user *User) error
	FindUserById(id string) (*User, error)
	FindAllUsers() ([]*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *User) error {
	user.Created = time.Now()
	user.ID = uuid.New().String()
	return s.repo.Create(user)
}

func (s *userService) FindUserById(id string) (*User, error) {
	return s.repo.FindById(id)
}

func (s *userService) FindAllUsers() ([]*User, error) {
	return s.repo.FindAll()
}
