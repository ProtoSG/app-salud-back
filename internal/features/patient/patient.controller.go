package patient

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (this *Controller) Create(w http.ResponseWriter, r *http.Request) {
	payloadPatient := &Patient{}
	if err := utils.ReadJSON(r, &payloadPatient); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := middleware.ValidateStruct(payloadPatient); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
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

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"id":      patient_id,
		"message": "Paciente creado",
	})
}

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
	disease := r.URL.Query().Get("disease")

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
		Disease:  disease,
	}

	patients, err := this.service.Read(page, limit, *filters)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, patients)
}
