package medicalhistory

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

// @Summary     Registra un nuevo antecedente médico
// @Description Crea un nuevo registro de historial médico para un paciente.
// @Tags        medicalhistory
// @Accept      json
// @Produce     json
// @Param       body  body      MedicalHistory       true  "Datos del antecedente médico"
// @Success     201   {object}  utils.Response                          "Antecedente creado exitosamente"
// @Failure     400   {object}  utils.ErrorResponse                     "Solicitud inválida"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /medicalhistory [post]
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payloadHistory := &MedicalHistory{}
	if err := utils.ReadJSON(r, &payloadHistory); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validatesErrors := middleware.ValidateStruct(payloadHistory); validatesErrors != nil {
		msg := strings.Join(validatesErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	id, err := c.service.Create(
		payloadHistory.PatientID,
		payloadHistory.Description,
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      id,
		Message: "Antecedente creado exitosamente",
	})
}

// @Summary     Obtiene antecedentes médicos por paciente
// @Description Recupera todos los registros de historial médico de un paciente.
// @Tags        medicalhistory
// @Produce     json
// @Param       id    path      int                   true  "ID del paciente"  example(123)
// @Success     200   {array}   MedicalHistoryBase                         "Lista de antecedentes médicos"
// @Failure     400   {object}  utils.ErrorResponse                     "Formato de ID inválido"
// @Failure     404   {object}  utils.ErrorResponse                     "No se encontraron antecedentes"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /medicalhistory/patient/{id} [get]
func (c *Controller) GetByPatientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "formato de ID incorrecto")
		return
	}

	histories, err := c.service.ReadByPatientID(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, histories)
}
