package patient

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (this *Controller) Create(w http.ResponseWriter, r *http.Request) {
	payloadPatient := &Patient{}
	if err := utils.ReadJSON(r, &payloadPatient); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := middleware.ValidateStruct(payloadPatient); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	patient_id, err := this.service.Create(
		payloadPatient.FirstName,
		payloadPatient.LastName,
		payloadPatient.DNI,
		payloadPatient.BirthDate,
		payloadPatient.Gender,
		payloadPatient.Address,
		payloadPatient.Phone,
		payloadPatient.Email,
		payloadPatient.PhotoURL,
	)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"id":      patient_id,
		"message": "Paciente creado",
	})
}

func (this *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	patients, err := this.service.Read()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, patients)
}
