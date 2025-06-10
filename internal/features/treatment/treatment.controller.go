package treatment

import (
	"net/http"
	"strconv"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
	"github.com/gorilla/mux"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (c *Controller) CreateTreatment(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en el contexto.")
		return
	}

	roleName := claims["role_name"].(string)
	if roleName != "DOCTOR" && roleName != "ENFERMERO" {
		utils.WriteError(w, http.StatusUnauthorized, "No estas autorizado.")
		return
	}

	payloadTreatment := &Treatment{}
	if err := utils.ReadJSON(r, &payloadTreatment); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payloadTreatment); validationErrors != nil {
		utils.WriteError(w, http.StatusBadRequest, validationErrors)
		return
	}

	treatmentID, err := c.service.Create(payloadTreatment)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message":      "Tratamiento creado correctamente",
		"treatment_id": treatmentID,
	})
}

func (c *Controller) GetTreatmentsByPatientID(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en el contexto.")
		return
	}

	roleName := claims["role_name"].(string)
	if roleName != "DOCTOR" && roleName != "ENFERMERO" {
		utils.WriteError(w, http.StatusUnauthorized, "No estas autorizado.")
		return
	}

	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "ID de paciente inválido.")
		return
	}

	treatments, err := c.service.ReadByPatientID(patientID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, treatments)
}
