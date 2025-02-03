package repository

import (
	"database/sql"
	"errors"
	"golang_starter_template/pkg/config/database"
	"golang_starter_template/pkg/jobs/entity"
	"golang_starter_template/pkg/utils"
)

type UserRepoImpl struct {
	db database.Database
}

// NewUserRepoImpl will create a new UserRepoImpl
func NewUserRepoImpl(db database.Database) *UserRepoImpl {
	return &UserRepoImpl{db}
}

// Create a new user
func (ur *UserRepoImpl) Create(user *entity.User) error {
	_, err := ur.db.GetDB().Exec(`INSERT INTO users (firstname, lastname, email, password, username, date_of_birth, phone, gender, avatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, user.FirstName, user.LastName, user.Email, user.Password, user.Username, user.DateOfBirth, user.Phone, user.Gender, user.Avatar)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

// GetByEmail will return user by email
func (ur *UserRepoImpl) GetByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := ur.db.GetDB().QueryRow(`SELECT * FROM users WHERE email = $1`, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Username, &user.DateOfBirth, &user.Phone, &user.Gender, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LoggerWarn.Println(utils.Warn + "No user found" + utils.Reset)
			return nil, nil // No user found
		}
		return nil, errors.New("failed to get user")
	}

	return user, nil
}

// GetByID will return user by id
func (ur *UserRepoImpl) GetByID(id int) (*entity.User, error) {
	user := new(entity.User)
	err := ur.db.GetDB().QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Username, &user.DateOfBirth, &user.Phone, &user.Gender, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LoggerWarn.Println(utils.Warn + "No user found" + utils.Reset)
			return nil, nil // No user found
		}
		return nil, errors.New("failed to get user")
	}

	return user, nil
}

// GetPassword will return user password
func (ur *UserRepoImpl) GetPassword(email string) (string, error) {
	var password string
	err := ur.db.GetDB().QueryRow(`SELECT password FROM users WHERE email = $1`, email).Scan(&password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LoggerWarn.Println(utils.Warn + "No user found" + utils.Reset)
			return "", nil // No user found
		}
		return "", errors.New("failed to get password")
	}

	return password, nil
}

// GetRoles will return roles
func (ur *UserRepoImpl) GetRoles() ([]entity.Role, error) {
	var roles []entity.Role
	rows, err := ur.db.GetDB().Query(`SELECT * FROM roles`)
	if err != nil {
		return nil, errors.New("failed to get roles")
	}

	for rows.Next() {
		var role entity.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, errors.New("failed to get roles")
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// GetRoleID will return role id
func (ur *UserRepoImpl) GetRoleID(role string) (int, error) {
	var id int
	err := ur.db.GetDB().QueryRow(`SELECT id FROM roles WHERE name = $1`, role).Scan(&id)
	if err != nil {
		return 0, errors.New("failed to get role id")
	}
	return id, nil
}

// AssignRole will assign role to user
func (ur *UserRepoImpl) AssignRole(userID int, roleID int) error {
	_, err := ur.db.GetDB().Exec(`INSERT INTO role_user (user_id, role_id) VALUES ($1, $2)`, userID, roleID)
	if err != nil {
		return errors.New("failed to assign role")
	}
	return nil
}

// GetRolesByUserID will return roles by user id
func (ur *UserRepoImpl) GetRolesByUserID(userID int) ([]entity.Role, error) {
	var roles []entity.Role
	rows, err := ur.db.GetDB().Query(`SELECT roles.id, roles.name FROM roles JOIN role_user ON roles.id = role_user.role_id WHERE role_user.user_id = $1`, userID)
	if err != nil {
		return nil, errors.New("failed to get roles")
	}

	for rows.Next() {
		var role entity.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, errors.New("failed to get roles")
		}
		roles = append(roles, role)
	}
}
