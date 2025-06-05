package patient

import "time"

type Patient struct {
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	DNI       string    `json:"dni" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	PhotoURL  string    `json:"photo_url" validate:"required"`
}

type PatientBasicData struct {
	PatientID string `json:"patient_id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}
