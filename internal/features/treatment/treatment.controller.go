package treatment

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

// @Summary     Crea un nuevo tratamiento
// @Description Registra un tratamiento para un paciente (solo DOCTOR o ENFERMERO).
// @Tags        treatment
// @Accept      json
// @Produce     json
// @Param       body  body      Treatment            true  "Datos del tratamiento"
// @Success     201   {object}  utils.Response                         "Tratamiento creado correctamente"
// @Failure     400   {object}  utils.ErrorResponse                    "Solicitud inválida"
// @Failure     401   {object}  utils.ErrorResponse                    "No autorizado"
// @Failure     500   {object}  utils.ErrorResponse                    "Error interno del servidor"
// @Router      /treatment [post]
func (c *Controller) CreateTreatment(w http.ResponseWriter, r *http.Request) {
	payloadTreatment := &Treatment{}
	if err := utils.ReadJSON(r, &payloadTreatment); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payloadTreatment); validationErrors != nil {
		msg := strings.Join(validationErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	treatmentID, err := c.service.Create(
		payloadTreatment.PatientID,
		payloadTreatment.DoctorID,
		payloadTreatment.StartDate,
		payloadTreatment.EndDate,
		payloadTreatment.Description,
		payloadTreatment.Observations,
	)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      treatmentID,
		Message: "Tratamiento creado correctamente",
	})
}

// @Summary     Obtiene tratamientos por ID de paciente
// @Description Recupera todos los tratamientos asociados a un paciente (solo DOCTOR o ENFERMERO).
// @Tags        treatment
// @Produce     json
// @Param       id    path      int                  true  "ID del paciente"
// @Success     200   {array}   TreatmentBase                             "Lista de tratamientos"
// @Failure     400   {object}  utils.ErrorResponse                   "ID inválido"
// @Failure     401   {object}  utils.ErrorResponse                   "No autorizado"
// @Failure     404   {object}  utils.ErrorResponse                   "Tratamientos no encontrados"
// @Failure     500   {object}  utils.ErrorResponse                   "Error interno del servidor"
// @Router      /treatment/patient/{id} [get]
func (c *Controller) GetTreatmentsByPatientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "ID de paciente inválido.")
		return
	}

	treatments, err := c.service.ReadByPatientID(patientID)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, treatments)
}
