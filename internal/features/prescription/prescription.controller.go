package prescription

import (
	"net/http"
	"strings"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

// @Summary     Registra una nueva receta médica
// @Description Crea una nueva prescripción para un paciente con sus ítems asociados.
// @Tags        prescription
// @Accept      json
// @Produce     json
// @Param       body  body      Prescription          true  "Datos de la receta"
// @Success     201   {object}  utils.Response                         "Receta creada exitosamente"
// @Failure     400   {object}  utils.ErrorResponse                    "Solicitud inválida"
// @Failure     500   {object}  utils.ErrorResponse                    "Error interno del servidor"
// @Router      /prescription [post]
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en contexto")
		return
	}

	doctorID := int(claims["user_id"].(float64))

	req := &Prescription{}
	if err := utils.ReadJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if validateErrors := middleware.ValidateStruct(req); validateErrors != nil {
		msg := strings.Join(validateErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	pres := &Prescription{
		PatientID:           req.PatientID,
		ElectronicSignature: req.ElectronicSignature,
		Observations:        req.Observations,
		Items:               req.Items,
	}

	prescriptionID, err := this.service.Create(doctorID, pres)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      prescriptionID,
		Message: "Receta creada.",
	})
}

// @Summary     Lista todas las recetas
// @Description Recupera todas las prescripciones almacenadas.
// @Tags        prescription
// @Produce     json
// @Success     200   {array}   PrescriptionBase                          "Lista de recetas"
// @Failure     500   {object}  utils.ErrorResponse                    "Error interno del servidor"
// @Router      /prescription [get]
func (this *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := this.service.Read()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, prescriptions)
}
