package labresult

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
	sampleDate time.Time,
	testType, resultValue, observations string,
) (int, error) {
	labResult := &LabResult{
		PatientID:    patientID,
		DoctorID:     doctorID,
		SampleDate:   sampleDate,
		TestType:     testType,
		ResultValue:  resultValue,
		Observations: observations,
	}

	return s.repo.Create(labResult)
}

func (s *Service) ReadByPatientID(id int) ([]*LabResultBase, error) {
	labResults, err := s.repo.ReadByPatientID(id)
	if err != nil {
		return nil, err
	}

	if len(labResults) == 0 {
		return nil, utils.NewEntityNotFound(id, "ID", "Lab Result")
	}

	return labResults, nil
}
