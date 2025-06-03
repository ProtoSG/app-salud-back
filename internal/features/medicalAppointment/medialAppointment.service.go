package medicalappointment

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) ReadAppointmentsToday(doctor_id int) ([]*MedicalAppointmentByDoctor, error) {
	return this.repo.ReadAppointmentsToday(doctor_id)
}
