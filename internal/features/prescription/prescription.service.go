package prescription

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) Create(pres *Prescription) (int, error) {
	newID, err := this.repo.Create(pres)
	if err != nil {
		return 0, fmt.Errorf("error al crear prescription en repo: %w", err)
	}

	return newID, nil
}

func (this *Service) Read() ([]*PrescriptionBase, error) {
	return this.repo.Read()
}
