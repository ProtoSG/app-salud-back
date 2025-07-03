package patient

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

// @Summary     Crea un nuevo paciente
// @Description Registra un nuevo paciente en la base de datos.
// @Tags        patient
// @Accept      json
// @Produce     json
// @Param       body      body      Patient         true  "Datos del paciente"
// @Success     201       {object}  utils.Response                   "Paciente creado exitosamente"
// @Failure     400       {object}  utils.ErrorResponse              "Solicitud inválida"
// @Failure     500       {object}  utils.ErrorResponse              "Error interno del servidor"
// @Router      /patient [post]
func (this *Controller) Create(w http.ResponseWriter, r *http.Request) {
	payloadPatient := &Patient{}
	if err := utils.ReadJSON(r, &payloadPatient); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validateErrors := middleware.ValidateStruct(payloadPatient); validateErrors != nil {
		msg := strings.Join(validateErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	patient_id, err := this.service.Create(
		payloadPatient.FirstName,
		payloadPatient.LastName,
		payloadPatient.DNI,
		payloadPatient.BirthDate,
		payloadPatient.Gender,
		payloadPatient.Address,
		payloadPatient.Phone,
		payloadPatient.Email,
		payloadPatient.PhotoURL,
	)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Response{
		ID:      patient_id,
		Message: "Paciente creado",
	})
}

// @Summary     Lista pacientes con filtros y paginación
// @Description Obtiene la lista de pacientes, con parámetros opcionales de página, límite, género, enfermedad y rango de edad.
// @Tags        patient
// @Produce     json
// @Param       page     query     int              false  "Número de página"        default(1)
// @Param       limit    query     int              false  "Resultados por página"   default(9)
// @Param       gender   query     string           false  "Filtrar por género"
// @Param       minAge   query     int              false  "Edad mínima"             default(0)
// @Param       maxAge   query     int              false  "Edad máxima"             default(0)
// @Success     200      {array}   PatientBasicData                            "Lista de pacientes"
// @Failure     400      {object}  utils.ErrorResponse               "Parámetros inválidos"
// @Failure     500      {object}  utils.ErrorResponse               "Error interno del servidor"
// @Router      /patient [get]
func (this *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	parseInt := func(key string, defaultVal int, errMsg string) (int, bool) {
		s := r.URL.Query().Get(key)
		if s == "" {
			return defaultVal, true
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errMsg)
			return 0, false
		}
		return v, true
	}

	page, ok := parseInt("page", 1, "Número de página inválido")
	if !ok {
		return
	}
	limit, ok := parseInt("limit", 9, "Número de límite inválido")
	if !ok {
		return
	}

	gender := r.URL.Query().Get("gender")

	minAge, ok := parseInt("minAge", 0, "Edad mínima inválida")
	if !ok {
		return
	}
	maxAge, ok := parseInt("maxAge", 0, "Edad máxima inválida")
	if !ok {
		return
	}
	if minAge > 0 && maxAge > 0 && minAge > maxAge {
		utils.WriteError(w, http.StatusBadRequest, "Edad mínima no puede ser mayor que edad máxima")
		return
	}

	filters := &PatientFilters{
		Gender:   gender,
		RangeAge: [2]int{minAge, maxAge},
	}

	patients, err := this.service.Read(page, limit, *filters)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, patients)
}

// @Summary     Obtiene un paciente por ID
// @Description Recupera el detalle de un paciente dado su identificador.
// @Tags        patient
// @Produce     json
// @Param       id    path      int              true  "ID del paciente"
// @Success     200   {object}  PatientInfo                         "Datos del paciente"
// @Failure     400   {object}  utils.ErrorResponse             "ID inválido"
// @Failure     404   {object}  utils.ErrorResponse             "Paciente no encontrado"
// @Failure     500   {object}  utils.ErrorResponse             "Error interno del servidor"
// @Router      /patient/{id} [get]
func (this *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Formato de ID inválido")
		return
	}

	patient, err := this.service.ReadByID(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, patient)
}
