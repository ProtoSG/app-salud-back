package auth

import (
	"fmt"
	"net/http"
	"strings"

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

// @Summary      Registra un nuevo usuario
// @Description  Crea un nuevo usuario en la base de datos.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      user.User            true  "Datos del usuario"
// @Success      201   {object}  utils.Response                              "Usuario creado correctamente"
// @Failure      400   {object}  utils.ErrorResponse                         "Solicitud inválida"
// @Failure      500   {object}  utils.ErrorResponse                         "Error interno del servidor"
// @Router       /register [post]
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	payload := &user.User{}
	if err := utils.ReadJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErrors := middleware.ValidateStruct(payload); validationErrors != nil {
		msg := strings.Join(validationErrors, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
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

	res := utils.Response{
		ID:      id,
		Message: "Usuario creado correctamente",
	}

	utils.WriteJSON(w, http.StatusCreated, res)
}

// @Summary     Iniciar sesión con un usuario
// @Description Autentica al usuario con email y contraseña, y establece una cookie con el JWT.
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       body  body      AuthLogin             true  "Credenciales del usuario"
// @Success     200   {object}  utils.AuthResponse   "Usuario autenticado correctamente"
// @Failure     400   {object}  utils.ErrorResponse   "Solicitud inválida"
// @Failure     401   {object}  utils.ErrorResponse   "Credenciales inválidas o usuario no encontrado"
// @Failure     500   {object}  utils.ErrorResponse   "Error interno del servidor"
// @Router      /login [post]
func (this *Controller) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := &AuthLogin{}
	if err := utils.ReadJSON(r, userLogin); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if details := middleware.ValidateStruct(userLogin); details != nil {
		msg := strings.Join(details, "; ")
		utils.WriteError(w, http.StatusBadRequest, msg)
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

	token, err := utils.CreateJWT(user.UserID, user.RolName, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	})

	utils.WriteJSON(w, http.StatusOK, utils.AuthResponse{
		ID:    user.UserID,
		Email: user.Email,
	})
}

func (this *Controller) Validate(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "No hay claims en contexto")
		return
	}

	id := int(claims["user_id"].(float64))
	roleName := claims["role_name"].(string)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"id":       id,
		"role_ame": roleName,
	})
}
