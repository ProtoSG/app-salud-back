package diagnosis

import (
	"log"
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

func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
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

	payloadDiagnosis := &Diagnosis{}
	if err := utils.ReadJSON(r, &payloadDiagnosis); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payloadDiagnosis); validationErrors != nil {
		utils.WriteError(w, http.StatusBadRequest, validationErrors)
		return
	}

	diagnosisID, err := this.service.Create(
		payloadDiagnosis.PatientID,
		payloadDiagnosis.DoctorID,
		payloadDiagnosis.DiagnosisDate,
		payloadDiagnosis.Description,
		payloadDiagnosis.Observations,
	)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"ID":      diagnosisID,
		"message": "Diagnostico creado exitosamente",
	})
}

func (this *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Formato ID inválido")
		return
	}

	diagnosis, err := this.service.ReadByID(id)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, diagnosis)
}
