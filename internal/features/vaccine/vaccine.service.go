package vaccine

import (
	"time"

	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(
	patientID int,
	vaccineType string,
	administeredOn time.Time,
	dose string,
	observations string,
) (int, error) {
	vaccine := &Vaccine{
		PatientID:      patientID,
		VaccineType:    vaccineType,
		AdministeredOn: administeredOn,
		Dose:           dose,
		Observations:   observations,
	}

	return s.repo.Create(vaccine)
}

func (s *Service) ReadByPatientID(id int) ([]*VaccineBase, error) {
	vaccines, err := s.repo.ReadByPatientID(id)
	if err != nil {
		return nil, err
	}

	if len(vaccines) == 0 {
		return nil, utils.NewEntityNotFound(id, "ID", "Vacuna")
	}

	return vaccines, nil
}
