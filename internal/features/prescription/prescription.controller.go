package prescription

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

func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	req := &Prescription{}
	if err := utils.ReadJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if validateErrors := middleware.ValidateStruct(req); validateErrors != nil {
		utils.WriteError(w, http.StatusBadRequest, validateErrors)
		return
	}

	pres := &Prescription{
		PatientID:           req.PatientID,
		DoctorID:            req.DoctorID,
		ElectronicSignature: req.ElectronicSignature,
		Observations:        req.Observations,
		Items:               req.Items,
	}

	prescriptionID, err := this.service.Create(pres)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"id":      prescriptionID,
		"message": "Receta creada.",
	})
}

func (this *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := this.service.Read()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, prescriptions)
}
