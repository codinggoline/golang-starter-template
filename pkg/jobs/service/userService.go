package service

import "golang_starter_template/pkg/jobs/entity"

type UserService interface {
	// Create a new user
	CreateUser(user *entity.User) error
}
