package treatment

type Service interface {
	Create(treatment *Treatment) (int, error)
	ReadByPatientID(id int) ([]*TreatmentBase, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(treatment *Treatment) (int, error) {
	return s.repo.Create(treatment)
}

func (s *service) ReadByPatientID(id int) ([]*TreatmentBase, error) {
	return s.repo.ReadByPatientID(id)
}
