package usuario

type User struct {
	RoleID    int    `json:"role_id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewUsuario(roleID int, firstName, lastName, email, password string) *User {
	return &User{
		RoleID:    roleID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
