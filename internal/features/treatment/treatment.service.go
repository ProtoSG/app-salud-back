package treatment

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
	patientID, doctorID int,
	startDate, endDate time.Time,
	description, observations string,
) (int, error) {
	treatment := &Treatment{
		PatientID:    patientID,
		DoctorID:     doctorID,
		StartDate:    startDate,
		EndDate:      endDate,
		Description:  description,
		Observations: observations,
	}
	return s.repo.Create(treatment)
}

func (s *Service) ReadByPatientID(id int) ([]*TreatmentBase, error) {
	treatments, err := s.repo.ReadByPatientID(id)
	if err != nil {
		return nil, err
	}

	if len(treatments) == 0 {
		return nil, utils.NewEntityNotFound(id, "ID", "Tratamiento")
	}

	return treatments, nil
}
