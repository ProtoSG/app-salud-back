package patient

import (
	"fmt"
	"time"

	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) Create(
	firstName,
	lastName,
	dni string,
	birthDate time.Time,
	gender,
	address,
	phone,
	email,
	photoUrl string,
) (int64, error) {
	if err := this.repo.FindPatientByDNI(dni); err == nil {
		return 0, fmt.Errorf("Paciente con DNI %s ya existe", dni)
	} else {
		if _, ok := err.(*utils.EntityNotFound); !ok {
			return 0, fmt.Errorf("Service Create: %w", err)
		}
	}

	patient := &Patient{
		FirstName: firstName,
		LastName:  lastName,
		DNI:       dni,
		BirthDate: birthDate,
		Gender:    gender,
		Address:   address,
		Phone:     phone,
		Email:     email,
		PhotoURL:  photoUrl,
	}

	return this.repo.Create(patient)
}

func (this *Service) Read() ([]*PatientBasicData, error) {
	return this.repo.Read()
}
