package interfaces

import "golang_starter_template/pkg/jobs/entity"

type UserRepo interface {
	// Create a new user
	Create(user *entity.User) error
	// GetByEmail will return user by email
	GetByEmail(email string) (*entity.User, error)
	// GetByID will return user by id
	GetByID(id int) (*entity.User, error)
	// GetPassword will return user password
	GetPassword(email string) (string, error)
}
