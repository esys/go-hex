package domain

type UserRepository interface {
	Create(user *User) error
	FindById(id string) (*User, error)
	FindAll() ([]*User, error)
	Save(user *User) error
}
