package user

type User struct {
	RoleID    int    `json:"role_id" validate:"required" example:"1"`
	FirstName string `json:"first_name" validate:"required" example:"Diego Alberto"`
	LastName  string `json:"last_name" validate:"required" example:"Salazar Garcia"`
	Email     string `json:"email" validate:"required,email" example:"diego@gmail.com"`
	Password  string `json:"password" validate:"required" example:"tuPasswordSecret"`
}

func NewUser(roleID int, firstName, lastName, email, password string) *User {
	return &User{
		RoleID:    roleID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
