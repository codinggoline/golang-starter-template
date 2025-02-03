package entity

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"firstname" db:"firstname"`
	LastName    string `json:"lastname" db:"lastname"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Username    string `json:"username" db:"username"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Phone       string `json:"phone" db:"phone"`
	Gender      string `json:"gender" db:"gender"`
	Avatar      string `json:"avatar" db:"avatar"`
	Roles       []Role `json:"roles" db:"-"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type Role struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
