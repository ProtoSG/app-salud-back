package labresult

import (
	"net/http"
	"strconv"
	"strings"

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

// @Summary     Registra un nuevo resultado de laboratorio
// @Description Crea un nuevo registro de resultado de laboratorio para un paciente (solo DOCTOR o ENFERMERO).
// @Tags        labresult
// @Accept      json
// @Produce     json
// @Param       body  body      LabResult             true  "Datos del resultado de laboratorio"
// @Success     201   {object}  utils.Response                           "Resultado de laboratorio creado exitosamente"
// @Failure     400   {object}  utils.ErrorResponse                      "Solicitud inválida"
// @Failure     401   {object}  utils.ErrorResponse                      "No autorizado"
// @Failure     500   {object}  utils.ErrorResponse                      "Error interno del servidor"
// @Router      /labresult [post]
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payloadLabResult := &LabResult{}
	if err := utils.ReadJSON(r, &payloadLabResult); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	validatesErrors := middleware.ValidateStruct(payloadLabResult)
	if validatesErrors != nil {
		msg := strings.Join(validatesErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
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

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      labResultID,
		Message: "resultado de laboratorio creado exitosamente",
	})
}

// @Summary     Obtiene resultados de laboratorio por ID de paciente
// @Description Recupera todos los resultados de laboratorio asociados a un paciente (solo DOCTOR o ENFERMERO).
// @Tags        labresult
// @Produce     json
// @Param       id    path      int                   true  "ID del paciente"
// @Success     200   {array}   LabResultBase                               "Lista de resultados de laboratorio"
// @Failure     400   {object}  utils.ErrorResponse                      "ID inválido"
// @Failure     401   {object}  utils.ErrorResponse                      "No autorizado"
// @Failure     404   {object}  utils.ErrorResponse                      "No se encontraron resultados"
// @Failure     500   {object}  utils.ErrorResponse                      "Error interno del servidor"
// @Router      /labresult/patient/{id} [get]
func (c *Controller) GetByPatientID(w http.ResponseWriter, r *http.Request) {
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
