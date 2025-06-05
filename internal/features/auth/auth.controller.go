package auth

import (
	"fmt"
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/features/user"
	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payload := &user.User{}
	if err := utils.ReadJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payload); validationErrors != nil {
		utils.WriteError(w, http.StatusBadRequest, validationErrors)
		return
	}

	id, err := this.service.CreateUser(
		payload.RoleID,
		payload.FirstName,
		payload.LastName,
		payload.Email,
		payload.Password,
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"id":      fmt.Sprintf("%d", id),
		"message": fmt.Sprintf("Usuario con ID %d creado", id),
	}

	utils.WriteJSON(w, http.StatusCreated, data)
}

func (this *Controller) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := &AuthLogin{}
	if err := utils.ReadJSON(r, userLogin); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if details := middleware.ValidateStruct(userLogin); details != nil {
		utils.WriteError(w, http.StatusBadRequest, details)
		return
	}

	user, err := this.service.Authenticate(userLogin.Email, userLogin.Password)
	if err != nil {
		switch err.(type) {
		case *utils.EntityNotFound:
			utils.WriteError(w, http.StatusUnauthorized, err.Error())
		case *utils.ErrInvalidCredentials:
			utils.WriteError(w, http.StatusUnauthorized, err.Error())
		default:
			utils.WriteError(w, http.StatusUnauthorized, fmt.Sprintf("Error interno: %v", err))
		}
		return
	}

	token, err := utils.CreateJWT(user.UserID, user.RolName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	})

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"id":    user.UserID,
		"email": user.Email,
	})
}

func (this *Controller) Validate(w http.ResponseWriter, r *http.Request) {
}
