package usuario

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) FindUser(email string) (bool, error) {
	found, err := this.repo.FindUser(email)
	if err != nil {
		return false, err
	}

	return found, nil
}
