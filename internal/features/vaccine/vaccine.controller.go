package vaccine

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

// @Summary     Registra un nuevo registro de vacuna
// @Description Crea un nuevo registro de vacunación para un paciente.
// @Tags        vaccine
// @Accept      json
// @Produce     json
// @Param       body  body      Vaccine              true  "Datos de la vacuna"
// @Success     201   {object}  utils.Response                          "Vacuna creada exitosamente"
// @Failure     400   {object}  utils.ErrorResponse                     "Solicitud inválida"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /vaccine [post]
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payloadVaccine := &Vaccine{}
	if err := utils.ReadJSON(r, &payloadVaccine); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validatesErros := middleware.ValidateStruct(payloadVaccine); validatesErros != nil {
		msg := strings.Join(validatesErros, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	id, err := c.service.Create(
		payloadVaccine.PatientID,
		payloadVaccine.VaccineType,
		payloadVaccine.AdministeredOn,
		payloadVaccine.Dose,
		payloadVaccine.Observations,
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      id,
		Message: "Vaccine creado exitosamente",
	})
}

// @Summary     Obtiene registros de vacuna por paciente
// @Description Recupera todos los registros de vacunación asociados a un paciente.
// @Tags        vaccine
// @Produce     json
// @Param       id    path      int                   true  "ID del paciente"  example(123)
// @Success     200   {array}   VaccineBase                                 "Lista de vacunas"
// @Failure     400   {object}  utils.ErrorResponse                     "Formato de ID inválido"
// @Failure     404   {object}  utils.ErrorResponse                     "No se encontraron registros de vacuna"
// @Failure     500   {object}  utils.ErrorResponse                     "Error interno del servidor"
// @Router      /vaccine/patient/{id} [get]
func (c *Controller) GetByPatientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "foramto de ID incorrecto")
		return
	}

	vaccines, err := c.service.ReadByPatientID(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, vaccines)
}
