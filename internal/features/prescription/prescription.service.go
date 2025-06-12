package prescription

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) Create(doctorID int, pres *Prescription) (int, error) {
	newID, err := this.repo.Create(doctorID, pres)
	if err != nil {
		return 0, fmt.Errorf("error al crear prescription en repo: %w", err)
	}

	return newID, nil
}

func (this *Service) Read() ([]*PrescriptionBase, error) {
	return this.repo.Read()
}
