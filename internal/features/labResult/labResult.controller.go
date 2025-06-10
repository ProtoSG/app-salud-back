package labresult

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

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en el contexto")
		return
	}

	roleName := claims["role_name"].(string)
	if roleName != "DOCTOR" && roleName != "ENFERMERO" {
		utils.WriteError(w, http.StatusUnauthorized, "No estas autorizado")
		return
	}

	payloadLabResult := &LabResult{}
	if err := utils.ReadJSON(r, &payloadLabResult); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	validatesErrors := middleware.ValidateStruct(payloadLabResult)
	if validatesErrors != nil {
		utils.WriteError(w, http.StatusBadRequest, validatesErrors)
		return
	}

	labResultID, err := c.service.Create(
		payloadLabResult.PatientID,
		payloadLabResult.DoctorID,
		payloadLabResult.SampleDate,
		payloadLabResult.TestType,
		payloadLabResult.ResultValue,
		payloadLabResult.Observations,
	)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"ID":      labResultID,
		"message": "resultado de laboratorio creado exitosamente",
	})
}

func (c *Controller) GetByPatientID(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en el contexto")
		return
	}

	roleName := claims["role_name"].(string)
	if roleName != "DOCTOR" && roleName != "ENFERMERO" {
		utils.WriteError(w, http.StatusUnauthorized, "No estas autorizado")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Formato ID inválido.")
		return
	}

	results, err := c.service.ReadByPatientID(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, results)
}
