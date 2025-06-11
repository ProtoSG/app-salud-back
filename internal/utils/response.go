package utils

type Response struct {
	ID      int    `json:"id" example:"1"`
	Message string `json:"message" example:"usuario creado exitosamente"`
}

type AuthResponse struct {
	ID    int    `json:"id" example:"1"`
	Email string `json:"email" example:"diego@gmail.com"`
}

type ErrorResponse struct {
	Status int    `json:"status" example:"400"`
	Error  string `json:"error"  example:"Error en la Solicitud"`
}
