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
	PatientID int    `json:"patient_id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}

type PatientFilters struct {
	Gender   string
	RangeAge [2]int
	Disease  string
}

type PatientInfo struct {
	PatientID int    `json:"patient_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Allergy   string `json:"allergy"`
}
