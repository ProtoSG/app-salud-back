package medicalappointment

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
