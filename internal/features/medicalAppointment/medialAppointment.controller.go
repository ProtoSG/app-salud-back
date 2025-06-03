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

func (this *Controller) GetAppointmentsToday(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en contexto")
		return
	}

	userID := int(claims["user_id"].(float64))
	// roleName := claims["roleName"].(string)

	appointments, err := this.service.ReadAppointmentsToday(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, appointments)
}
