package medicalappointment

import (
	"net/http"
	"strings"
	"time"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

// @Summary     Obtiene las citas del día
// @Description Retorna las citas médicas programadas para hoy para el usuario autenticado (solo DOCTOR o ENFERMERO).
// @Tags        medicalappointment
// @Produce     json
// @Success     200 {array}   MedicalAppointmentByDoctor     "Lista de citas para hoy"
// @Failure     401 {object}  utils.ErrorResponse   "No autorizado"
// @Failure     500 {object}  utils.ErrorResponse   "Error interno del servidor"
// @Router      /medicalappointment/today [get]
func (this *Controller) GetAppointmentsToday(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en contexto")
		return
	}

	userID := int(claims["user_id"].(float64))

	appointments, err := this.service.ReadAppointmentsToday(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, appointments)
}

// @Summary     Crea una nueva cita médica
// @Description Registra una cita médica para un paciente (requiere rol DOCTOR o ENFERMERO).
// @Tags        medicalappointment
// @Accept      json
// @Produce     json
// @Param       body  body      MedicalAppointment             true  "Datos de la cita médica"
// @Success     201   {object}  utils.Response                  "Cita médica creada correctamente"
// @Failure     400   {object}  utils.ErrorResponse             "Solicitud inválida"
// @Failure     401   {object}  utils.ErrorResponse             "No autorizado"
// @Failure     500   {object}  utils.ErrorResponse             "Error interno del servidor"
// @Router      /medicalappointment [post]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	payloadAppointment := &MedicalAppointment{}
	if err := utils.ReadJSON(r, &payloadAppointment); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validatesError := middleware.ValidateStruct(payloadAppointment); validatesError != nil {
		msg := strings.Join(validatesError, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	id, err := c.service.Create(
		payloadAppointment.PatientID,
		payloadAppointment.DoctorID,
		payloadAppointment.AppointmentTime,
		payloadAppointment.DurationMinutes,
		payloadAppointment.Reason,
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      id,
		Message: "Cita médica creada correctamente",
	})
}

// @Summary     Lista citas en un rango
// @Description Obtiene todas las citas del usuario autenticado entre dos fechas.
// @Tags        medicalappointment
// @Accept      json
// @Produce     json
// @Param       start  query     string   true  "Fecha de inicio ISO8601"   example(2025-04-06T00:00:00Z)
// @Param       end    query     string   true  "Fecha de fin ISO8601"      example(2025-04-12T23:59:59Z)
// @Success     200    {array}   MedicalAppointmentEvent    "Lista de eventos para el calendario"
// @Failure     400    {object}  utils.ErrorResponse        "Parámetros inválidos"
// @Failure     401    {object}  utils.ErrorResponse        "No autorizado"
// @Failure     500    {object}  utils.ErrorResponse        "Error interno del servidor"
// @Router      /medicalappointment [get]
func (c *Controller) ListRange(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")
	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "start inválido")
		return
	}
	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "end inválido")
		return
	}

	events, err := c.service.ReadByDateRange(start, end)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, events)
}
