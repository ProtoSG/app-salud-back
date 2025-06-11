package medicalhistory

import "github.com/ProtoSG/app-salud-back/internal/utils"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(patientID int, desctiption string) (int, error) {
	history := &MedicalHistory{
		PatientID:   patientID,
		Description: desctiption,
	}

	return s.repo.Create(history)
}

func (s *Service) ReadByPatientID(id int) ([]*MedicalHistoryBase, error) {
	histories, err := s.repo.ReadByPatientID(id)
	if err != nil {
		return nil, err
	}

	if len(histories) == 0 {
		return nil, utils.NewEntityNotFound(id, "ID", "Antecedente")
	}

	return histories, nil
}
