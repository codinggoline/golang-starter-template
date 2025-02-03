package impl

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"golang_starter_template/pkg/jobs/entity"
	"golang_starter_template/pkg/jobs/repository/interfaces"
	"strings"
)

type UserServiceImpl struct {
	Repository interfaces.UserRepo
}

// CreateUser will create a new user
func (us *UserServiceImpl) CreateUser(user *entity.User) error {
	// Check if required fields are not empty
	if strings.TrimSpace(user.FirstName) == "" || strings.TrimSpace(user.LastName) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.DateOfBirth) == "" || strings.TrimSpace(user.Phone) == "" || strings.TrimSpace(user.Gender) == "" {
		return errors.New("missing required field")
	}

	// Set default avatar based
	if strings.TrimSpace(user.Avatar) == "" {
		switch user.Gender {
		case "male":
			user.Avatar = "https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png"
		case "female":
			user.Avatar = "https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359554_960_720.png"
		default:
			return errors.New("invalid gender")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	return us.Repository.Create(user)
}

// GetUserByEmail will return user by email
