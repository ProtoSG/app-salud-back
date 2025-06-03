package main

import (
	"log"

	"github.com/ProtoSG/app-salud-back/cmd/api"
	"github.com/ProtoSG/app-salud-back/internal/config"
	"github.com/ProtoSG/app-salud-back/internal/db"
)

func main() {
	config := config.NewConfig()

	initDB := db.NewDBConnection(config.URL)
	db := initDB.SetupDB()

	app := api.NewAPIServer(":"+config.PORT, db)
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
