package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/router"
	"github.com/ProtoSG/app-salud-back/internal/services"
	"github.com/ProtoSG/app-salud-back/internal/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (this *APIServer) Run() error {
	serviceContainer := services.NewServiceContainer(this.db)

	r := mux.NewRouter()
	r.HandleFunc("/", this.checkHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()
	router.NewRouterContainer(apiRouter, serviceContainer)

	corsmiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
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
		Handler: corsmiddleware(r),
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
