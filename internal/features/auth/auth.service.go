package auth

import (
	"fmt"

	"github.com/ProtoSG/app-salud-back/internal/features/user"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) CreateUser(roleID int, firstName, lastName, email, password string) (int, error) {
	_, err := this.repo.FindUserByEmail(email)
	if err == nil {
		return 0, fmt.Errorf("usuario con email %s ya existe", email)
	}
	if _, ok := err.(*utils.EntityNotFound); !ok {
		return 0, fmt.Errorf("service Create: %w", err)
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return 0, fmt.Errorf("service Create: %w", err)
	}

	user := user.NewUser(roleID, firstName, lastName, email, hashedPassword)
	return this.repo.CreateUser(user)
}

func (this *Service) Authenticate(email, password string) (*UserAuth, error) {
	user, err := this.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(
		user.Password,
		password,
	) {
		return nil, utils.NewErrInvalidCredentials("Contraseña incorrecta")
	}

	return user, nil
}
