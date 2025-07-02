// @title App Salud API
// @version 1.0
// @description API REST para gestion de salud
// @BasePath /api
package api

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/ProtoSG/app-salud-back/docs"
	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/ProtoSG/app-salud-back/internal/router"
	"github.com/ProtoSG/app-salud-back/internal/services"
	"github.com/ProtoSG/app-salud-back/internal/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	addr      string
	db        *sql.DB
	originUrl string
}

func NewAPIServer(addr string, db *sql.DB, originUrl string) *APIServer {
	return &APIServer{
		addr:      addr,
		db:        db,
		originUrl: originUrl,
	}
}

func (this *APIServer) Run() error {
	serviceContainer := services.NewServiceContainer(this.db)

	r := mux.NewRouter()
	r.HandleFunc("/", this.checkHandler)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()
	router.NewRouterContainer(apiRouter, serviceContainer)

	corsmiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{this.originUrl}),
		handlers.AllowedMethods([]string{
			"get",
			"post",
			"put",
			"delete",
			"options",
		}),

		handlers.AllowedHeaders([]string{
			"content-type",
			"authorization",
		}),

		handlers.ExposedHeaders([]string{
			"authorization",
		}),

		handlers.AllowCredentials(),
	)

	svr := &http.Server{
		Addr:    this.addr,
		Handler: corsmiddleware(middleware.Logging(r)),
	}

	log.Printf("Listening on: http://localhost%s", this.addr)
	if err := svr.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (this *APIServer) checkHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Server active",
		"status":  "ok",
	}

	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error en el servidor")
		return
	}
}
