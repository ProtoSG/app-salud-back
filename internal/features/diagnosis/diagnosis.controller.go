package diagnosis

import (
	"log"
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

// @Summary     Registra un nuevo diagnóstico
// @Description Crea un nuevo diagnóstico para un paciente (solo DOCTOR o ENFERMERO).
// @Tags        diagnosis
// @Accept      json
// @Produce     json
// @Param       body  body      Diagnosis             true  "Datos del diagnóstico"
// @Success     201   {object}  utils.Response                          "Diagnóstico creado exitosamente"
// @Failure     400   {object}  utils.ErrorResponse                     "Solicitud inválida"
// @Failure     401   {object}  utils.ErrorResponse                     "No autorizado"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /diagnosis [post]
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payloadDiagnosis := &Diagnosis{}
	if err := utils.ReadJSON(r, &payloadDiagnosis); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payloadDiagnosis); validationErrors != nil {
		msg := strings.Join(validationErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
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

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      diagnosisID,
		Message: "Diagnostico creado exitosamente",
	})
}

// @Summary     Obtiene un diagnóstico por ID
// @Description Recupera un diagnóstico existente por su identificador (solo DOCTOR o ENFERMERO).
// @Tags        diagnosis
// @Produce     json
// @Param       id    path      int                   true  "ID del diagnóstico"
// @Success     200   {object}  DiagnosisBase                               "Detalle del diagnóstico"
// @Failure     400   {object}  utils.ErrorResponse                     "ID inválido"
// @Failure     401   {object}  utils.ErrorResponse                     "No autorizado"
// @Failure     404   {object}  utils.ErrorResponse                     "Diagnóstico no encontrado"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /diagnosis/{id} [get]
func (this *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
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
