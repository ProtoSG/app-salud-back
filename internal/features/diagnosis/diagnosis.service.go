package diagnosis

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

func (this *Service) Create(
	patientID, doctorID int, diagnosisDate time.Time,
	description, observations string,
) (int, error) {
	diagnosis := &Diagnosis{
		PatientID:     patientID,
		DoctorID:      doctorID,
		DiagnosisDate: diagnosisDate,
		Description:   description,
		Observations:  observations,
	}

	return this.repo.Create(diagnosis)
}

func (this *Service) ReadByID(id int) ([]*DiagnosisBase, error) {
	results, err := this.repo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, utils.NewEntityNotFound(id, "ID", "Diagnosis")
	}

	return results, nil
}
